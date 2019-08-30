package model

import (
	"time"
)

const (
	RoleStateOff = 0
	RoleStateOn  = 1
)

var (
	RoleAdmin = Role{
		Id:         99,
		Name:       "管理员",
		State:      RoleStateOn,
		CreateTime: time.Now(),
		EditTime:   time.Now(),
	}
)

type Role struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT:false"`
	Name        string `gorm:"type:varchar(50)"`
	State       int16  `gorm:"type:tinyint"`
	Description string `gorm:"type:varchar(500)"`
	CreatUserId int    `gorm:"type:int"`
	EditTime    time.Time
	CreateTime  time.Time
	Users       []*User    `gorm:"many2many:auth_user_role_relation;"`
	Services    []*Service `gorm:"many2many:auth_service_role_relation;"`
}

// 通过TableName方法将User表命名为`s_userinfo`
func (Role) TableName() string {
	return "auth_user_role"
}
