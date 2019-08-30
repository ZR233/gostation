/*
@Time : 2019-06-26 15:45
@Author : zr
@File : user
@Software: GoLand
*/
package model

import (
	"camdig/server/errors"
	"time"
	"unicode/utf8"
)

type UserStatus int16

const (
	UserStatusBan    UserStatus = 0
	UserStatusNormal UserStatus = 1
)

type User struct {
	Id         int        `gorm:"primary_key;AUTO_INCREMENT"`
	Usercode   string     `gorm:"type:varchar(20)"`
	Email      string     `gorm:"type:varchar(64)"`
	Phone      string     `gorm:"type:varchar(20)"`
	WXOpenId   string     `gorm:"type:varchar(60)"`
	Password   string     `gorm:"type:char(64)"`
	Name       string     `gorm:"type:varchar(20)"`
	Status     UserStatus `gorm:"type:tinyint;default:1"`
	Memo       string     `gorm:"type:varchar(200)"`
	CreateTime time.Time
	EditTime   time.Time
	LastOpt    time.Time
	Roles      []*Role `gorm:"many2many:auth_user_role_relation;Column:user_id"`

	UserDao UserDaoInterface `gorm:"-"` // 忽略本字段
}

func (u *User) SetPassword(password string) error {

	if utf8.RuneCountInString(password) < 6 {
		return errors.ErrPasswordFmt
	}

	u.Password = password
	return nil
}

// 通过TableName方法将User表命名为`s_user`
func (User) TableName() string {
	return "s_user"
}
func (u User) CheckStatus() error {
	switch u.Status {
	case UserStatusBan:
		return errors.ErrUserBan
	case UserStatusNormal:
		return nil
	default:
		return errors.ErrUserStatus
	}
}

func (u *User) GetRoles() (roles []*Role) {
	return u.UserDao.GetRoles(u)
}
