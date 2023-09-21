package router

import "github.com/gin-gonic/gin"

// SetRouter 配置路由
func SetRouter() *gin.Engine {
	var r = gin.Default()

	User(r)

	return r
}
