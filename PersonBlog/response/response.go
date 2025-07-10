package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` //返回数据 可选
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: nil})
}
