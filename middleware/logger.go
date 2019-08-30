/*
@Time : 2019-06-28 17:48
@Author : zr
@File : logger
@Software: GoLand
*/
package middleware

import (
	"camdig/server/errors"
	"camdig/server/lib/log"
	"camdig/server/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginTime := time.Now()

		defer func() {
			code := 0
			memo := ""

			if err := recover(); err != nil {
				err_, ok := err.(error)
				if ok {
					memo = err_.Error()
					code = errors.UNKNOWN
				}
			}

			userid := 0
			src := ""

			endTime := time.Now()

			s, exists := c.Get("session")
			if exists {
				s, ok := s.(*model.Session)
				if ok {
					userid = s.UserId
					src = s.Channel
				}
			}

			err_, exists := c.Get("err")
			if exists {
				err2, ok := err_.(errors.StandardError)
				if ok {
					code = err2.Code()
					memo = err2.DebugMsg()
				}
			} else {
				code = c.GetInt("code")
				memo = c.GetString("memo")
			}

			postForm := c.Request.PostForm
			params, _ := json.Marshal(postForm)

			url := c.Request.URL.Path
			log.Logger.WithFields(logrus.Fields{
				"execTime":  endTime.Sub(beginTime) / time.Millisecond,
				"apiName":   url,
				"optTime":   beginTime,
				"optUserid": userid,
				"src":       src,
				"code":      code,
				"params":    params,
			}).Info(memo)

		}()

		c.Next()
	}
}
