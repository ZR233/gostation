/*
@Time : 2019-07-5 15:45
@Author : zr
@Software: GoLand
*/
package model

import (
	"time"
)

type ApiLog struct {
	OptTime   time.Time `gorm:"index:idx_api_log_opt_time"`
	ApiName   string    `gorm:"type:varchar(200)"`
	Code      int       `gorm:"type:int"`
	ExecTime  int       `gorm:"type:int"`
	OptUserid int       `gorm:"type:int"`
	Src       string    `gorm:"type:varchar(20)"`
	Params    string    `gorm:"type:varchar(200)"`
	Msg       string    `gorm:"type:varchar(200)"`
	Trace     string    `gorm:"type:varchar(500)"`
}

// 通过TableName方法将User表命名为`s_userinfo`
func (ApiLog) TableName() string {
	return "s_log_api"
}
