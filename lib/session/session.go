/*
@Time : 2019-08-21 11:19
@Author : zr
*/
package session

import (
	"camdig/server/lib/redis"
	"github.com/ZR233/session"
	"github.com/spf13/viper"
	"sync"
)

var sessionManager *session.Manager
var once sync.Once

func newSM() {
	v := viper.GetViper()
	profix := v.GetString("redis.prefix")

	db := session.NewRedisAdapter(redis.GetRedis(), profix+"_session")
	sessionManager = session.NewManager(db)
}

func GetSessionManager() *session.Manager {
	once.Do(newSM)

	return sessionManager
}
