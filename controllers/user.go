package controllers

import (
	"daily_plan_go/common/api"
	"daily_plan_go/services/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCtl struct {
	api.Api
}

func (c UserCtl) HelloWorld(ctx *gin.Context) {
	req := new(dto.Hello)
	err := c.MakeContext(ctx).Bind(&req).Errors
	if err != nil {
		msg := fmt.Sprintf("参数绑定失败，err：%s", err.Error())
		c.BadRequest(msg)
		return
	}

	c.Logger.Info("执行user逻辑")

	msg := fmt.Sprintf("%s say: 'Hello world!'", req.Name)
	c.Ok(msg)
}
