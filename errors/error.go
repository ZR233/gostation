/*
@Time : 2019-06-26 16:21
@Author : zr
@File : error
@Software: GoLand
*/
package errors

import "fmt"

const (
	UNKNOWN             = -1
	SUCCESS             = 0
	Busy                = 1
	ServiceNotAvailable = 2
	KeyNotExist         = 3
	ZipErr              = 4
	NamePwdErr          = 1001
	UserExists          = 1002
	PicVerifyCodeErr    = 1003
	PermissionDenied    = 1004
	PasswordFmtErr      = 1005
	PasswordErr         = 1006
	UserBan             = 1007
	UserStatusErr       = 1008
	TokenErr            = 1010
	SendTooFast         = 1011
	UserNotExists       = 1012
	VFCodeErr           = 1013
	NoRecord            = 1014
	ParamErr            = 5000
	DATABASE            = 9000
	REDIS               = 10000
	HDFS                = 11000
)

var (
	ErrSuccess             = NewServiceErr(SUCCESS, "执行成功")
	ErrBusy                = NewServiceErr(Busy, "服务器繁忙，请稍后再试")
	ErrServiceNotAvailable = NewServiceErr(ServiceNotAvailable, "服务暂不可用")
	ErrKeyNotExist         = NewInternalErr(KeyNotExist, "key不存在")
	ErrPermissionDenied    = NewServiceErr(PermissionDenied, "权限不足")
	ErrNamePwdIncorrect    = NewServiceErr(NamePwdErr, "用户名或密码错误")
	ErrPicVerifyCode       = NewServiceErr(PicVerifyCodeErr, "图形验证码错误")
	ErrPasswordFmt         = NewServiceErr(PasswordFmtErr, "密码格式错误")
	ErrPassword            = NewServiceErr(PasswordErr, "密码错误")
	ErrUserExists          = NewServiceErr(UserExists, "用户已存在")
	ErrSendTooFast         = NewServiceErr(SendTooFast, "发送频率过高")
	ErrUserBan             = NewServiceErr(UserBan, "用户被禁用")
	ErrUserStatus          = NewServiceErr(UserBan, "用户状态异常")
	ErrUserNotExists       = NewServiceErr(UserNotExists, "用户不存在")
	ErrNoRecord            = NewServiceErr(NoRecord, "记录不存在")
	ErrVFCode              = NewServiceErr(VFCodeErr, "验证码错误")
	ErrToken               = NewServiceErr(TokenErr, "登录失效")
)

type StdErrorInterface interface {
	Code() int
	DebugMsg() string
	ShowMsg() string
}

//StandardError 标准错误，包含错误码和错误信息
type StandardError struct {
	code     int
	debugMsg string
	showMsg  string
}

func (s StandardError) Code() int {
	return s.code
}

func (s StandardError) DebugMsg() string {
	return s.debugMsg
}

func (s StandardError) ShowMsg() string {
	return s.showMsg
}

func NewInternalErr(code int, msg string) StandardError {
	return StandardError{
		code:     code,
		debugMsg: msg,
		showMsg:  "服务器内部错误",
	}
}
func NewServiceErr(code int, msg string) StandardError {
	return StandardError{
		code:     code,
		debugMsg: msg,
		showMsg:  msg,
	}
}
func NewParamErr(paramName string) StandardError {
	return StandardError{
		code:     ParamErr,
		debugMsg: "参数错误:" + paramName,
		showMsg:  "参数错误",
	}
}
func RedisError(err error) error {
	if err != nil {
		return NewInternalErr(REDIS, err.Error())
	} else {
		return nil
	}
}
func DatabaseError(err error) error {
	if err != nil {
		return NewInternalErr(DATABASE, err.Error())
	} else {
		return nil
	}
}
func HDFSError(err error) error {
	if err != nil {
		return NewInternalErr(HDFS, err.Error())
	} else {
		return nil
	}
}

// Error 实现了 Error接口
func (s StandardError) Error() string {
	return fmt.Sprintf("errorCode: %d, errorMsg %s", s.code, s.debugMsg)
}

func FromError(err error) StdErrorInterface {
	if err != nil {
		err2, ok := err.(StandardError)
		if ok {
			return err2
		} else {
			r := StandardError{
				code:     UNKNOWN,
				debugMsg: err.Error(),
				showMsg:  "内部错误",
			}
			return r
		}
	} else {
		return nil
	}

}
