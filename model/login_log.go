/*
@Time : 2019-07-16 15:09
@Author : zr
*/
package model

import "time"

type LoginLog struct {
	Id         int64
	UserId     int
	Ip         string `gorm:"type:varchar(60)"`
	CreateTime time.Time
	Result     int
	Src        string `gorm:"type:varchar(40)"`
}

func (LoginLog) TableName() string {
	return "s_log_login"
}
