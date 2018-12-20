package models

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//AuthorizationModel authorized identity
type AuthorizationModel struct {
	Id uint `form:"id" json:"id,omitempty" comment:""`

	Username string `form:"username" json:"username,omitempty" comment:"昵称/登陆用户名"`

	FullName string `form:"full_name" json:"full_name,omitempty" comment:"真实姓名"`

	Email string `form:"email" json:"email,omitempty" comment:"邮箱"`

	Mobile string `form:"mobile" json:"mobile,omitempty" comment:"手机号码"`

	Password string `form:"password" json:"password,omitempty" comment:"密码"`

	RoleId uint `form:"role_id" json:"role_id,omitempty" comment:"角色ID:2-超级用户,4-普通用户"`

	Status uint `form:"status" json:"status,omitempty" comment:"状态: 1-正常,2-禁用/删除"`

	Avatar string `form:"avatar" json:"avatar,omitempty" comment:"用户头像"`

	Remark string `form:"remark" json:"remark,omitempty" comment:"备注"`

	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" comment:""`

	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" comment:""`
}

//TableName auth table name
func (m *AuthorizationModel) TableName() string {
	return "users"
}

//One get one
func (m *AuthorizationModel) One() (one *AuthorizationModel, err error) {
	one = &AuthorizationModel{}
	err = crudOne(m, one)
	return
}

//All get all
func (m *AuthorizationModel) All(q *PaginationQuery) (list *[]AuthorizationModel, total uint, err error) {
	list = &[]AuthorizationModel{}
	total, err = crudAll(m, q, list)
	return
}

//Update a row
func (m *AuthorizationModel) Update() (err error) {
	where := AuthorizationModel{Id: m.Id}
	m.Id = 0
	m.makePassword()
	return crudUpdate(m, where)
}
func (m *AuthorizationModel) makePassword() {
	if m.Password != "" {
		if bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost); err != nil {
			logrus.WithError(err).Error("bcrypt making password is failed")
		} else {
			m.Password = string(bytes)
		}
	}
}

//Create insert a row
func (m *AuthorizationModel) Create() (err error) {
	m.Id = 0
	m.makePassword()
	return mysqlDB.Create(m).Error
}

//Delete a row
func (m *AuthorizationModel) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

//Login by user and record by IP
func (m *AuthorizationModel) Login(ip string) (*jwtObj, error) {
	m.Id = 0
	if m.Password == "" {
		return nil, errors.New("password is required")
	}
	inputPassword := m.Password
	m.Password = ""
	loginTryKey := "login:" + ip
	loginRetries, _ := mem.GetUint(loginTryKey)
	if loginRetries > uint(viper.GetInt("app.login_try")) {
		memExpire := viper.GetInt("app.mem_expire_min")
		return nil, fmt.Errorf("for too many wrong login retries the %s will ban for login in %d minitues", ip, memExpire)
	}
	//you can implement more detailed login retry rule
	//for i don't know what your login username i can't implement the ip+username rule in my boilerplate project
	// about username and ip retry rule

	err := mysqlDB.Where(m).First(&m).Error
	if err != nil {
		//username fail ip retries add 5
		loginRetries = loginRetries + 5
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
	//password is set to bcrypt check
	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(inputPassword)); err != nil {
		// when password failed reties will add 1
		loginRetries = loginRetries + 1
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
	key := fmt.Sprintf("login:%d", m.Id)
	m.Password = ""
	//save login user  into the memory store
	data ,err := jwtGenerateToken(m)
	mem.Set(key, data)
	return data,err
}
