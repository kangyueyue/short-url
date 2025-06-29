package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
)

// @Summary Hello 示例接口
// @Description 这是一个Hello World接口
// @Tags 测试
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello World"
// @Router /hello [get]
func (s *HelloSvr) Hello(c *gin.Context) {
	http.Success(c, "Hello World")
}
