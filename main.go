/*
@Time : 2019-06-26 9:31
@Author : zr
@Software: GoLand
*/
package main

import (
	"camdig/server/config"
	"camdig/server/controller"
	"camdig/server/dao/gorm"
	"camdig/server/docs"
	"camdig/server/global"
	"camdig/server/lib/redis"
	"camdig/server/middleware"
	"camdig/server/router"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"os"
	"path"
)

// @title GoStation API 文档
// @version 1.0
// @description WifiDig 建站基本框架
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	config.Init("./")
	gorm.Init()
	redis.Init()

	if !global.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}

	//设置swagger文档主机
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", global.ServerHost(), global.ServerPort())

	// Logging to a file.
	logFile := path.Join(global.LogPath(), "gin_err.log")
	f, _ := os.Create(logFile)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := router.NewDefaultRouter()

	r.Use(middleware.Logger())
	r.Use(middleware.Session())

	//自动绑定controller类中所有公有方法,所有方法都为post方式
	r.AutoRegisterController("/v1", &controller.Auth{})
	r.AutoRegisterController("/v1", &controller.User{})
	//debug模式下启用swag文档
	if global.Debug() {
		// 注册swagger文档
		// 查看地址为 /v1/SignIn/docs/index.html     /v1/SignIn/docs/doc.json
		r.GET("/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// listen and serve on 0.0.0.0:8080
	_ = r.Run(fmt.Sprintf("%s:%s", global.ServerListen(), global.ServerPort()))
}
