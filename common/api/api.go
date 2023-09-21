package api

import (
	"daily_plan_go/common/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Api struct {
	Context *gin.Context
	Logger  *logrus.Logger
	Errors  error
}

func (a *Api) addError(err error) {
	if a.Errors == nil {
		a.Errors = err
	} else if err != nil {
		a.Errors = fmt.Errorf("%w", err)
	}
}

// MakeContext 设置http上下文
func (a *Api) MakeContext(c *gin.Context) *Api {
	a.Context = c
	a.Logger, _ = logger.InitLogger()
	a.apiInfoLog()
	return a
}

func (a *Api) apiInfoLog() {
	a.Logger.WithFields(logrus.Fields{
		"user_id":    "luguanzong",
		"ip":         a.Context.Request.RemoteAddr,
		"query":      "query",
		"request_id": "ec2bf8e55a11474392f8867e92624e04",
	}).Info("请求信息")
}

// Bind 绑定参数
func (a *Api) Bind(req interface{}) *Api {
	m := a.Context.Request.Method
	if m != "POST" && m != "GET" {
		a.addError(fmt.Errorf("无法识别请求类型, 仅支持get和post方法"))
		return a
	}

	var err error
	switch m {
	case "POST":
		err = a.Context.ShouldBindJSON(req)
	case "GET":
		err = a.Context.ShouldBindQuery(req)
	default:
		err = a.Context.ShouldBindUri(req)
	}
	if err != nil {
		a.addError(err)
	}

	return a
}

type JsonReturn struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (a *Api) Success(msg string, data interface{}) {
	json := &JsonReturn{Code: http.StatusOK, Msg: msg, Data: data}
	a.Context.JSON(http.StatusOK, json)
}

func (a *Api) Ok(msg string) {
	json := &JsonReturn{Code: http.StatusOK, Msg: msg}
	a.Context.JSON(http.StatusOK, json)
}

func (a *Api) BadRequest(msg string) {
	json := &JsonReturn{Code: http.StatusBadRequest, Msg: msg}
	a.Context.JSON(http.StatusBadRequest, json)
}

func (a *Api) Error(msg string) {
	json := &JsonReturn{Code: http.StatusInternalServerError, Msg: msg}
	a.Context.JSON(http.StatusInternalServerError, json)
}
