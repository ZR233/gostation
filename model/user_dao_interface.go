/*
@Time : 2019-08-21 9:35
@Author : zr
*/
package model

type UserDaoInterface interface {
	GetRoles(user *User) []*Role
}
