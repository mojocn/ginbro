package models

import (
	"errors"
	"time"
	{{if .IsAuthTable}}"fmt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"github.com/sirupsen/logrus"
	{{end}}
)

var _ = time.Thursday
//{{.ModelName}}
type {{.ModelName}} struct {
	{{range .Columns}}
	{{.ModelTag}}{{end}}
}
//TableName
func (m *{{.ModelName}}) TableName() string {
	return "{{.TableName}}"
}
//One
func (m *{{.ModelName}}) One() (one *{{.ModelName}}, err error) {
	one = &{{.ModelName}}{}
	err = crudOne(m, one)
	return
}
//All
func (m *{{.ModelName}}) All(q *PaginationQuery) (list *[]{{.ModelName}}, total uint, err error) {
	list = &[]{{.ModelName}}{}
	total, err = crudAll(m, q, list)
	return
}
//Update
func (m *{{.ModelName}}) Update() (err error) {
	where := {{.ModelName}}{Id: m.Id}
	m.Id = 0
	{{if .IsAuthTable }}m.makePassword()
	{{end}}
	return crudUpdate(m, where)
}
//Create
func (m *{{.ModelName}}) Create() (err error) {
	m.Id = 0
    {{if .IsAuthTable }}m.makePassword()
    {{end}}
	return mysqlDB.Create(m).Error
}
//Delete
func (m *{{.ModelName}}) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
{{if .IsAuthTable }}

//Login
func (m *{{.ModelName}}) Login(ip string) (*jwtObj, error) {
	m.Id = 0
	if m.{{.PasswordPropertyName}} == "" {
		return nil, errors.New("password is required")
	}
	inputPassword := m.{{.PasswordPropertyName}}
	m.{{.PasswordPropertyName}} = ""
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
	if err := bcrypt.CompareHashAndPassword([]byte(m.{{.PasswordPropertyName}}), []byte(inputPassword)); err != nil {
		// when password failed reties will add 1
		loginRetries = loginRetries + 1
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
    m.{{.PasswordPropertyName}} = ""
	key := fmt.Sprintf("login:%d", m.Id)

	//save login user  into the memory store

    data ,err := jwtGenerateToken(m)
    mem.Set(key, data)
    return data,err
}

func (m *{{.ModelName}}) makePassword() {
	if m.{{.PasswordPropertyName}} != "" {
		if bytes, err := bcrypt.GenerateFromPassword([]byte(m.{{.PasswordPropertyName}}), bcrypt.DefaultCost); err != nil {
			logrus.WithError(err).Error("bcrypt making password is failed")
		} else {
			m.{{.PasswordPropertyName}} = string(bytes)
		}
	}
}

{{end}}