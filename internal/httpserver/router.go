package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterConfig 路由配置
type RouterConfig struct {
	Method  string
	Handler func(ctx *gin.Context)
}

// GetRounter 获取路由
func (s *Server) GetRounter() map[string]*RouterConfig {
	return map[string]*RouterConfig{
		"/hello": {
			Method:  http.MethodGet,
			Handler: s.HelloSvr.Hello,
		},
	}
}
