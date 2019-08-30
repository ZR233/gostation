/*
@Time : 2019-06-26 9:31
@Author : zr
@Software: GoLand
*/
package controller

import (
	"camdig/server/service"
	"net/http"
)

type Auth struct {
	BaseController
}

// SignOut godoc
// @Summary 用户退出
// @tags Auth
// @Description 用户退出
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param token formData string true "Token"
// @Success 200 {string} success
// @Router /Auth/SignOut [post]
func (a *Auth) SignOut() {
	session := a.getSession()
	s := service.NewAuthService()
	err := s.SignOut(session)
	r := a.handleErr(err)
	if err != nil {
		return
	}
	a.ctx.JSON(http.StatusOK, r)
}

type LoginResponse struct {
	ResponseBase
	Rs map[string]string
}

// SignUp godoc
// @Summary 用户注册
// @tags Auth
// @Description 所有用户用此接口注册
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param userCode formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string false "邮箱"
// @Param phone formData string false "手机号"
// @Param name formData string false "姓名"
// @Success 200 {object} controller.ResponseBase
// @Router /Auth/SignUp [post]
func (a *Auth) SignUp() {
	userCode := a.ctx.PostForm("userCode")
	password := a.ctx.PostForm("password")
	email := a.ctx.PostForm("email")
	phone := a.ctx.PostForm("phone")
	name := a.ctx.PostForm("name")

	s := service.NewAuthService()
	_, err := s.SignUp(userCode, password, email, phone, name)
	rb := a.handleErr(err)
	if err != nil {
		return
	}

	a.ctx.JSON(http.StatusOK, rb)
}

type ResponseVerifyPic struct {
	Id  string
	Pic string
}

// GenVerifyPic godoc
// @Summary 获取验证图片
// @tags Auth
// @Description 获取验证图片和验证id
// @Accept  x-www-form-urlencoded
// @Produce json
// @Success 200 {object} controller.ResponseVerifyPic
// @Router /Auth/GenVerifyPic [post]
func (a *Auth) GenVerifyPic() {

	s := service.NewAuthService()
	id, pic := s.GenVerifyPic()

	r := ResponseVerifyPic{
		id,
		pic,
	}

	a.ctx.JSON(http.StatusOK, r)
}

// SignIn godoc
// @Summary 用户登录
// @tags Auth
// @Description 用户登录并返回token
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param verifyName formData string true "用户名,邮箱,手机号"
// @Param verifyCode formData string true "密码"
// @Param src formData string true "来源"
// @Param picId formData string true "图形验证码id"
// @Param picCode formData string true "图形验证码"
// @Success 200 {object} controller.LoginResponse
// @Router /Auth/SignIn [post]
func (a *Auth) SignIn() {
	verifyName := a.ctx.PostForm("verifyName")
	verifyCode := a.ctx.PostForm("verifyCode")
	src := a.ctx.PostForm("src")
	picId := a.ctx.PostForm("picId")
	picCode := a.ctx.PostForm("picCode")
	ip := a.getIp()

	s := service.NewAuthService()
	session, err := s.SignIn(verifyName, verifyCode, src, picId, picCode, ip)
	a.handleSignIn(session, err)
}
