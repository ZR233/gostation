/*
@Time : 2019-07-09 10:54
@Author : zr
*/
package test

import (
	"camdig/server/config"
	"camdig/server/dao/gorm"
	"camdig/server/global"
	"camdig/server/lib/redis"
	"camdig/server/utils"
)

func init() {
	dirPath := utils.GetSrcRoot()
	config.Init(dirPath)
	global.Test = true
	gorm.Init()
	redis.Init()
}
