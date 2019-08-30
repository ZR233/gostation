/*
@Time : 2019-06-26 16:06
@Author : zr
@File : Auth
@Software: GoLand
*/
package service

import (
	"camdig/server/dao/gorm"
	"camdig/server/errors"
	"camdig/server/lib/session"
	"camdig/server/model"
	"github.com/mojocn/base64Captcha"
	"strconv"
	"time"
)

type Auth struct {
	base
}

func NewAuthService() *Auth {
	s := &Auth{
		base: *newBase(""),
	}
	return s
}
func saveLoginLog(err error, ip string, userId int, src string) {
	loginLogDao := gorm.GetLoginLogDAO()
	code := -1
	if err != nil {
		err_ := errors.FromError(err)
		code = err_.Code()
	} else {
		code = errors.SUCCESS
	}

	err = loginLogDao.Save(ip, code, userId, src)
	if err != nil {
		panic(err)
	}
}
func (a *Auth) SignUp(userCode string, password string, email string, phone string, name string) (*model.User, error) {
	userDAO := a.getUserDao()
	return userDAO.CreateUser(userCode, password, email, phone, name, "")
}

func (a *Auth) SignIn(verifyName string, password string, src string, picId string, picCode string, ip string) (s *model.Session, err error) {
	userId := 0
	defer saveLoginLog(err, ip, userId, src)

	//验证图形验证码
	verifyResult := base64Captcha.VerifyCaptcha(picId, picCode)
	if !verifyResult {
		err = errors.ErrPicVerifyCode
		return
	}
	userDAO := a.getUserDao()
	user, err := userDAO.GetUserByVerifyName(verifyName)
	if err != nil {
		return
	}
	userId = user.Id
	err = user.CheckStatus()
	if err != nil {
		return
	}
	if !userDAO.PasswordCorrect(user.Id, password) {
		err = errors.ErrNamePwdIncorrect
		return
	}

	sm := session.GetSessionManager()
	s_, err := sm.CreateSession(strconv.Itoa(userId), src, time.Now().Add(time.Hour*24*5))
	if err != nil {
		return
	}
	s = model.NewSession(s_)

	return
}
func (a *Auth) SignOut(s *model.Session) (err error) {
	sm := a.getSessionManager()
	sess, _ := sm.FindByToken(s.Token)
	err = sm.Delete(sess)
	return
}

func (a *Auth) UserInfo(sess *model.Session) (err error) {
	_, err = a.getUserByToken(sess)
	return
}
func (a *Auth) GenVerifyPic() (id string, pic string) {

	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 4,
	}
	id, capA := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	pic = base64Captcha.CaptchaWriteToBase64Encoding(capA)

	return id, pic
}
