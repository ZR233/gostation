/*
@Time : 2019-06-26 17:14
@Author : zr
@File : redis
@Software: GoLand
*/
package redis

import (
	_ "camdig/server/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
	"sync"
)

var client *redis.Client
var mu sync.Mutex

func GetRedis() *redis.Client {
	return client
}
func connect() {

	redisConf := viper.GetStringMapString("redis")
	db, err := strconv.Atoi(redisConf["db"])
	if err != nil {
		panic(err)
	}

	conn := redis.NewClient(&redis.Options{
		Addr:     redisConf["host"],
		Password: redisConf["password"], // no password set
		DB:       db,                    // use default DB
	})
	client = conn
	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func ReConnect() error {
	mu.Lock()
	defer mu.Unlock()

	if err := client.Close(); err != nil {
		return err
	}
	connect()
	return nil
}

func Init() {
	connect()
}
