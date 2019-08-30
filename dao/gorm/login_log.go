/*
@Time : 2019-07-08 10:57
@Author : zr
*/
package gorm

import (
	"camdig/server/errors"
	"camdig/server/model"
	"time"
)

type LoginLogDAO struct {
	*BaseSqlDAO
}

func GetLoginLogDAO() LoginLogDAO {
	m := LoginLogDAO{
		BaseSqlDAO: newBase(),
	}
	return m
}

func (l LoginLogDAO) Save(ip string, result int, userId int, src string) error {

	loginlog := &model.LoginLog{
		Ip:         ip,
		Result:     result,
		CreateTime: time.Now(),
		UserId:     userId,
		Src:        src,
	}
	err := l.getDB().Save(loginlog).Error
	return errors.DatabaseError(err)
}
