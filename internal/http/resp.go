package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功-200
func Success(c *gin.Context, data interface{}, msg ...string) {
	message := "操作成功"
	if len(msg) > 0 {
		message = strings.Join(msg, " ")
	}
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  message,
		Data: data,
	})
}

// Fail 失败-500
func Fail(c *gin.Context, msg ...string) {
	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  strings.Join(msg, " "),
		Data: nil,
	})
}
