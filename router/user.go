package router

import (
	"daily_plan_go/controllers"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	api := &controllers.UserCtl{}
	g := r.Group("user")
	{
		g.GET("/hello", api.HelloWorld)
	}
}
