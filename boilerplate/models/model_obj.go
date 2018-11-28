package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Department mysql table departments
type Department struct {
	Id             uint       `gorm:"column:id" form:"id" json:"id" comment:"部门(业务)表主键" sql:"int(10) unsigned,PRI"`
	ParentId       uint       `gorm:"column:parent_id" form:"parent_id" json:"parent_id" comment:"父级部门ID" sql:"int(10) unsigned,MUL"`
	Name           string     `gorm:"column:name" form:"name" json:"name" comment:"部门名称" sql:"varchar(50)"`
	OpsName        string     `gorm:"column:ops_name" form:"ops_name" json:"ops_name" comment:"运维工程师姓名" sql:"varchar(50)"`
	LeaderName     string     `gorm:"column:leader_name" form:"leader_name" json:"leader_name" comment:"负责人姓名" sql:"varchar(50)"`
	MemberIds      string     `gorm:"column:member_ids" form:"member_ids" json:"member_ids" comment:"must a string can unmarsh to an Object. 成员IDS json数组,关联users表" sql:"json"`
	EquipmentCount uint       `gorm:"column:equipment_count" form:"equipment_count" json:"equipment_count" comment:"" sql:"int(255) unsigned"`
	CreatedAt      *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
}

//TableName of GORM model
func (m *Department) TableName() string {
	return "departments"
}

//One find one row
func (m *Department) One() (one *Department, err error) {
	one = &Department{}
	err = crudOne(m, one)
	return
}

//All get all for pagination
func (m *Department) All(q *PaginationQuery) (list *[]Department, total uint, err error) {
	list = &[]Department{}
	total, err = crudAll(m, q, list)
	return
}

//Update a row
func (m *Department) Update() (err error) {
	where := Department{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

//Create insert a row
func (m *Department) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

//Delete destroy a row
func (m *Department) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
