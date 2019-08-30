/*
@Time : 2019-08-22 14:33
@Author : zr
*/
package middleware

import (
	"camdig/server/lib/session"
	"camdig/server/model"
	"github.com/ZR233/session/serr"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.PostForm("token")
		if token != "" {
			sm := session.GetSessionManager()
			sess, err := sm.FindByToken(token)
			if err != nil {
				if err != serr.TokenNotFound {
					panic(err)
				}
			} else {
				sess_ := model.NewSession(sess)
				c.Set("session", sess_)
			}
		}

		c.Next()
	}
}
