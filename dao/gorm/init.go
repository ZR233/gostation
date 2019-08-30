/*
@Time : 2019-08-19 17:55
@Author : zr
*/
package gorm

import (
	_ "camdig/server/config"
	"camdig/server/errors"
	"camdig/server/global"
	"camdig/server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sync"
)

var client *gorm.DB
var mu sync.Mutex
var config map[string]string

const DbConfName = "db1"
const DBTestConfName = "db_test"

func connect() {

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config["user"], config["password"], config["host"], config["port"], config["database_name"], config["charset"])

	var err error
	client, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	client.DB().SetMaxOpenConns(10)
	client.DB().SetMaxIdleConns(20)

	v := viper.GetViper()
	debug := v.GetBool("debug")
	client.LogMode(debug)
}
func ReConnect() error {
	mu.Lock()
	defer mu.Unlock()

	if err := client.Close(); err != nil {
		return err
	}
	connect()
	return nil
}

func setMigrate() {
	if err := client.AutoMigrate(
		&model.ApiLog{},
		&model.LoginLog{},
		&model.User{},
		&model.Service{},
		&model.Role{},
	).Error; err != nil {
		panic(err)
	}

	insertRoleDefault()
}

func insertRoleDefault() {
	//为管理员用户添加权限
	model.RoleAdmin.Services = append(model.RoleAdmin.Services, &model.ServiceUserManage)

	//增加管理员用户角色
	if err := client.Save(model.RoleAdmin).Error; err != nil {
		panic(err)
	}
}
func getConfig() map[string]string {
	v := viper.GetViper()
	var dbConf map[string]string
	if global.Test {
		dbConf = v.GetStringMapString("db." + DBTestConfName)
	} else {
		dbConf = v.GetStringMapString("db." + DbConfName)
	}
	return dbConf
}

func Init() {
	config = getConfig()
	connect()
	setMigrate()
}

func DbError(db *gorm.DB) error {
	if db.RecordNotFound() {
		return errors.ErrNoRecord
	}
	err := errors.DatabaseError(db.Error)
	if err != nil {
		panic(err)
	}
	return nil
}
