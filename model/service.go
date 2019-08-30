/*
@Time : 2019-07-5 15:45
@Author : zr
@Software: GoLand
*/
package model

import (
	"time"
)

var (
	ServiceUserManage = Service{
		Id:          1,
		SuperiorId:  0,
		Name:        "UserManage",
		Description: "用户管理",
		CreateTime:  time.Now(),
	}
	ServiceMsgManage = Service{
		Id:          2,
		SuperiorId:  0,
		Name:        "MsgManage",
		Description: "消息管理",
		CreateTime:  time.Now(),
	}
	ServiceStatisticManage = Service{
		Id:          3,
		SuperiorId:  0,
		Name:        "StatisticManage",
		Description: "管理端统计信息",
		CreateTime:  time.Now(),
	}
	ServiceSystemManage = Service{
		Id:          4,
		SuperiorId:  0,
		Name:        "SystemManage",
		Description: "系统管理",
		CreateTime:  time.Now(),
	}
)

type Service struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT:false"`
	SuperiorId  int    `gorm:"type:int"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(200)"`
	CreateTime  time.Time
	EditTime    time.Time
	Roles       []*Role `gorm:"many2many:auth_service_role_relation;"`
}

func (Service) TableName() string {
	return "auth_service"
}
