/*
@Time : 2019-07-16 14:36
@Author : zr
*/
package service

import (
	"camdig/server/dao/gorm"
	"camdig/server/errors"
	"camdig/server/lib/session"
	"camdig/server/model"
	session2 "github.com/ZR233/session"
	"runtime"
	"strings"
)

type base struct {
	restFulApiBase []string
	restFulApi     []string
	userId         int
}

func (b *base) GetUserId() int {
	return b.userId
}

func newBase(ApiBaseUrl string) (b *base) {
	b = &base{}
	if ApiBaseUrl == "" {
		return
	}

	b.restFulApiBase = strings.Split(ApiBaseUrl, "/")
	return
}
func (b *base) GetServiceTrace() []string {
	arr := append(b.restFulApiBase, b.restFulApi...)
	return arr
}
func (b *base) GetUrl() string {
	arr := append(b.restFulApiBase, b.restFulApi...)

	return strings.Join(arr, "/")
}

func (b *base) getCallers(skip int) {
	pc, _, _, _ := runtime.Caller(skip)
	name := runtime.FuncForPC(pc).Name()
	arr := strings.Split(name, "/")
	callersStr := arr[len(arr)-1]
	b.restFulApi = strings.Split(callersStr, ".")
	for i, v := range b.restFulApi {
		v = strings.TrimLeft(v, "(")
		v = strings.TrimLeft(v, "*")
		b.restFulApi[i] = strings.TrimRight(v, ")")
	}
	b.restFulApi = b.restFulApi[1:]
}

// 该函数会通过Session查找用户，并获取用户角色，检查用户是否拥有调用该函数的函数的权限
func (b *base) getUserByToken(session *model.Session) (user *model.User, err error) {

	if session == nil {
		err = errors.ErrToken
		return
	}

	b.userId = session.UserId
	b.getCallers(2)
	b.GetServiceTrace()

	user, err = gorm.NewUserDAO().GetUserById(b.userId)
	if err != nil {
		return
	}

	err = b.CheckUserIdServiceAuth(user)
	if err != nil {
		return
	}
	return
}

func (b *base) CheckUserIdServiceAuth(user *model.User) error {

	rolesUserHas := user.GetRoles()
	rolesNeed, err := b.GetRolesServiceNecessary(b.GetUrl())
	if err != nil {
		return err
	}

	for _, roleNeed := range rolesNeed {
		hasRole := false
		for _, roleHas := range rolesUserHas {
			if roleNeed.Id == roleHas.Id {
				hasRole = true
				break
			}
		}
		if !hasRole {
			return errors.ErrPermissionDenied
		}
	}

	return nil
}
func (b *base) GetRolesServiceNecessary(serviceName string) (roles []model.Role, err error) {

	serviceDao := gorm.GetServiceDAO()
	roles, err = serviceDao.GetRolesServiceNecessary(serviceName)
	return roles, err
}

func (b *base) getSessionManager() *session2.Manager {
	return session.GetSessionManager()
}

func (b *base) getUserDao() *gorm.UserDAO {
	return gorm.NewUserDAO()
}
