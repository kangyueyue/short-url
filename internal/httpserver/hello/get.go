package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
)

// Hello hello
func (s *HelloSvr) Hello(c *gin.Context) {
	http.Success(c, "Hello World")
}
