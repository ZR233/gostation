/*
@Time : 2019-08-22 16:58
@Author : zr
*/
package controller

import (
	"camdig/server/service"
	"net/http"
)

type User struct {
	BaseController
}

// SignOut godoc
// @Summary 用户信息
// @tags 用户
// @Description 用户信息
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param token formData string true "Token"
// @Success 200 {string} success
// @Router /User/UserInfo [post]
func (u *User) UserInfo() {
	sess := u.getSession()
	s := service.NewAuthService()
	err := s.UserInfo(sess)
	r := u.handleErr(err)
	if err != nil {
		return
	}
	u.ctx.JSON(http.StatusOK, r)
}
