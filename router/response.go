package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrCode int64

const (
	ERROR_FIELD             = 10001
	ERROR_WRONG_USER_OR_PWD = 10002
	ERROR_USER_NOT_FOUND    = 10003
	ERROR_LOGIN_ERROR       = 10004
	ERROR_REG_ERROR         = 10005
	ERROR_ALREADY_EXIST     = 10006
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func sendOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data: data,
	})
	c.Abort()
}
func sendFieldError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Code: ERROR_FIELD,
		Msg:  msg,
	})
}

func sendError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
	c.Abort()
}
