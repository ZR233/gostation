/*
@Time : 2019-06-26 9:31
@Author : zr
@File : router
@Software: GoLand
*/
package router

import (
	"camdig/server/controller"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

type Router struct {
	*gin.Engine
}

type Group struct {
	*gin.RouterGroup
}

func NewDefaultRouter() *Router {
	r := &Router{
		gin.Default(),
	}
	return r
}

/*
自动注册controller类的所有公有方法
路由地址为 relativePath/controller名/方法名
*/
func (r *Router) AutoRegisterController(relativePath string, c controller.Controller) *Router {
	autoRegisterController(r, true, relativePath, c)
	return r
}

type IRouter interface {
	POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}

/*
自动注册controller类的所有公有方法
路由地址为 relativePath/controller名/方法名
*/
func autoRegisterController(router IRouter, isPost bool, relativePath string, c interface{}) {
	typeOfController := reflect.TypeOf(c)
	name := typeOfController.String()
	namePath := strings.Split(name, ".")
	name = namePath[1]

	num := typeOfController.NumMethod()

	for i := 0; i < num; i++ {
		methodName := typeOfController.Method(i).Name
		finalPath := relativePath + "/" + name + "/" + methodName

		iter := i

		f := func(context *gin.Context) {
			controllerCase := reflect.New(typeOfController.Elem())
			baseController := controllerCase.Elem().FieldByName("BaseController")

			base := controller.NewBaseController(context)
			valueOfBase := reflect.ValueOf(*base)
			baseController.Set(valueOfBase)

			in := make([]reflect.Value, 0)
			//in[0] = valueOfContext

			controllerCase.Method(iter).Call(in)

		}
		if isPost {
			router.POST(finalPath, f)
		} else {
			router.GET(finalPath, f)
		}
	}
}

/*
自动注册controller类的所有公有方法
路由地址为 relativePath/controller名/方法名
*/
func (r *Group) AutoRegisterController(relativePath string, c controller.Controller) *Group {
	autoRegisterController(r, true, relativePath, c)
	return r
}

func (r *Router) Group(relativePath string, handlers ...gin.HandlerFunc) *Group {
	g := Group{}
	g.RouterGroup = r.Engine.Group(relativePath, handlers...)
	return &g
}
