/*
@Time : 2019-07-15 14:32
@Author : zr
*/
package middleware

import (
	"camdig/server/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type base struct {
	c *gin.Context
}

func (b base) handleErr(err error) {
	rs, err_ := controller.ResponseBaseFromError(err)
	b.c.Set("err", err_)
	b.c.Set("code", rs.Code)
	b.c.Set("memo", rs.Memo)
	b.c.JSON(http.StatusOK, rs)
	b.c.Abort()
	b.c.Next()
}

func newBase(c *gin.Context) base {
	b := base{
		c,
	}
	return b
}
