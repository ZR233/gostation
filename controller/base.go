/*
@Time : 2019-06-26 9:31
@Author : zr
@Software: GoLand
*/
package controller

import (
	"camdig/server/errors"
	"camdig/server/global"
	"camdig/server/model"
	"camdig/server/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Controller interface {
}
type BaseController struct {
	ctx *gin.Context
}

func NewBaseController(ctx *gin.Context) *BaseController {
	b := BaseController{
		ctx: ctx,
	}
	return &b
}

type ResponseBase struct {
	Code int
	Memo string
}

func ResponseBaseFromError(err error) (ResponseBase, errors.StdErrorInterface) {
	errStd := errors.FromError(err)
	if err == nil {
		errStd = errors.ErrSuccess
	}
	msg := ""
	if global.Debug() {
		msg = errStd.DebugMsg()
	} else {
		msg = errStd.ShowMsg()
	}

	return ResponseBase{
		Code: errStd.Code(),
		Memo: msg,
	}, errStd
}

func (b *BaseController) handleErr(err error) ResponseBase {
	r, err_ := ResponseBaseFromError(err)
	if err != nil {
		b.ctx.JSON(http.StatusOK, r)
	}
	b.ctx.Set("err", err_)
	b.ctx.Set("code", r.Code)
	b.ctx.Set("memo", r.Memo)
	return r
}

func (b *BaseController) postFromInt64(key string) (value int64, err error) {
	valueStr := b.ctx.PostForm(key)
	value, err = strconv.ParseInt(valueStr, 10, 64)
	b.handleErr(err)
	return value, err
}
func (b *BaseController) postFromInt(key string) (value int, err error) {
	valueStr := b.ctx.PostForm(key)
	value, err = strconv.Atoi(valueStr)
	b.handleErr(err)
	return value, err
}
func (b *BaseController) postFromFloat64(key string) (value float64, err error) {
	valueStr := b.ctx.PostForm(key)
	value, err = strconv.ParseFloat(valueStr, 64)
	b.handleErr(err)
	return value, err
}

func (b *BaseController) getPostFromTime(key string) (value time.Time, err error) {
	valueStr := b.ctx.PostForm(key)
	value, err = time.Parse(utils.TimeFormatter(), valueStr)
	b.handleErr(err)
	return value, err
}
func (b *BaseController) getPostFromDate(key string) (value time.Time, err error) {
	valueStr := b.ctx.PostForm(key)
	value, err = time.Parse(utils.TimeFormatterDate(), valueStr)
	b.handleErr(err)
	return value, err
}
func (b BaseController) getIp() string {
	// 获取tcp/ip
	addr := b.ctx.Request.RemoteAddr
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(addr)); err == nil {
		return ip
	}
	return ""
}
func (b BaseController) postFromNotNecessaryInt(key string) (*int, error) {
	intStr := b.ctx.PostForm(key)
	if intStr == "" {
		return nil, nil
	}
	value, err := strconv.Atoi(intStr)
	b.handleErr(err)
	return &value, err
}
func (b BaseController) postFromNotNecessaryDate(key string) (*time.Time, error) {
	str := b.ctx.PostForm(key)
	if str == "" {
		return nil, nil
	}
	value, err := time.Parse(utils.TimeFormatterDate(), str)
	b.handleErr(err)
	return &value, err
}
func (b BaseController) postFromNotNecessaryTime(key string) (*time.Time, error) {
	str := b.ctx.PostForm(key)
	if str == "" {
		return nil, nil
	}
	value, err := time.Parse(utils.TimeFormatter(), str)
	b.handleErr(err)
	return &value, err
}
func (b BaseController) formData(key string) (data []byte, err error) {
	file, err := b.ctx.FormFile(key)
	b.handleErr(err)
	if err != nil {
		return nil, errors.NewParamErr("formData")
	}
	src, err := file.Open()
	b.handleErr(err)
	if err != nil {
		return nil, errors.NewParamErr("formData")
	}
	data, err = ioutil.ReadAll(src)
	b.handleErr(err)
	if err != nil {
		return nil, errors.NewParamErr("formData")
	}
_:
	src.Close()

	return data, nil
}

func (b BaseController) formDataNotNecessary(key string) (data []byte, err error) {
	file, err := b.ctx.FormFile(key)
	if err != nil {
		if err.Error() == "http: no such file" {
			return data, errors.ErrKeyNotExist
		}
	}
	b.handleErr(err)
	if err != nil {
		return nil, errors.NewParamErr("formData")
	}
	src, err := file.Open()
	b.handleErr(err)
	if err != nil {
		return nil, errors.NewParamErr("formData")
	}
	data, err = ioutil.ReadAll(src)
	b.handleErr(err)
	if err != nil {
		return nil, errors.NewParamErr("formData")
	}
_:
	src.Close()

	return data, nil
}

func (b BaseController) formFile(key string) (file *multipart.FileHeader, err error) {
	file, err = b.ctx.FormFile(key)
	if err != nil {
		if err.Error() == "http: no such file" {
			err = errors.ErrKeyNotExist
		}
	}
	return
}

func (b BaseController) getToken() string {
	return b.ctx.PostForm("token")
}
func (b BaseController) getSession() *model.Session {
	if s, ok := b.ctx.Get("session"); ok {
		s, ok := s.(*model.Session)
		if ok {
			return s
		}
	}
	return nil
}
func (b BaseController) handleSignIn(session *model.Session, err error) {
	rb := b.handleErr(err)
	if err != nil {
		return
	}
	token := ""
	b.ctx.Set("session", session)
	token = session.Token

	r := LoginResponse{
		ResponseBase: rb,
		Rs: map[string]string{
			"token": token,
		},
	}

	b.ctx.JSON(http.StatusOK, r)
}
