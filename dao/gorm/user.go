package gorm

import (
	"camdig/server/errors"
	"camdig/server/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

/*
@Time : 2019-06-26 15:45
@Author : zr
@File : user
@Software: GoLand
*/

type UserDAO struct {
	*BaseSqlDAO
}

func NewUserDAO() *UserDAO {
	m := &UserDAO{
		BaseSqlDAO: newBase(),
	}
	return m
}

func (u UserDAO) NewUser() (user *model.User) {
	user = &model.User{}
	user.UserDao = &u
	return
}

func (u UserDAO) GetUserById(id int) (user *model.User, err error) {
	user = u.NewUser()
	user.Id = id
	err = u.getErr(u.getDB().Find(user))
	return
}

func (u *UserDAO) GetRoles(user *model.User) []*model.Role {
	var roles []*model.Role
	if err := u.getDB().Model(user).Related(&roles, "Roles").Error; err != nil {
		logrus.Info(err.Error())
	}

	var returnRoles []*model.Role
	for _, v := range roles {
		if v.State == model.RoleStateOn {
			returnRoles = append(returnRoles, v)
		}
	}
	return returnRoles
}
func (u *UserDAO) GetUserByVerifyName(verifyName string) (user *model.User, err error) {
	user = &model.User{}
	db := u.db.Where("usercode = ? or email = ? or phone = ?", verifyName, verifyName, verifyName).First(user)
	err = u.getErr(db)
	return
}

func encodePassword(password string) string {
	ctx := md5.New()
	password = password + ")_camDig"
	ctx.Write([]byte(password))
	return hex.EncodeToString(ctx.Sum(nil))
}

func (u *UserDAO) PasswordCorrect(id int, password string) (r bool) {
	r = false

	user := &model.User{}
_:
	u.getErr(u.db.Where("id = ?", id).First(user))

	password = encodePassword(password)
	r = password == user.Password
	return
}

func (u *UserDAO) CreateUser(
	userCode string,
	password string,
	email string,
	phone string,
	name string,
	memo string) (user *model.User, err error) {

	//若用户名，email，手机号都为空
	if userCode == "" && email == "" && phone == "" {
		err = errors.NewParamErr("user email phone 都为空")
		return
	}

	if password == "" {
		password = strconv.Itoa(rand.Int())
	}

	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = DbError(tx)
	if err != nil {
		panic(err)
	}

	tx = tx.Model(&model.User{})
	if userCode != "" {
		tx = tx.Or(model.User{Usercode: userCode})
	}
	if email != "" {
		tx = tx.Or(model.User{Email: email})
	}
	if phone != "" {
		tx = tx.Or(model.User{Phone: phone})
	}
	count := 0
	tx.Set("gorm:query_option", "FOR UPDATE").Count(&count)
	err = DbError(tx)
	if err != nil {
		panic(err)
	}

	if count > 0 {
		err = errors.ErrUserExists
		panic(err)
	}
	password = encodePassword(password)
	user = &model.User{
		Usercode:   userCode,
		Email:      email,
		Phone:      phone,
		Password:   password,
		Name:       name,
		Status:     model.RoleStateOn,
		Memo:       memo,
		CreateTime: time.Now(),
	}
_:
	DbError(tx.Create(user))

	return
}
