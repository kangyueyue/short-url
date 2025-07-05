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
		"short_url/set": {
			Method:  http.MethodPost,
			Handler: s.ShortUrlSvr.Set,
		},
		"short_url/get": {
			Method:  http.MethodPost,
			Handler: s.ShortUrlSvr.Get,
		},
		"short_url/list": {
			Method:  http.MethodGet,
			Handler: s.ShortUrlSvr.List,
		},
		"short_url/del": {
			Method:  http.MethodDelete,
			Handler: s.ShortUrlSvr.Del,
		},
		"ss/:id": {
			Method:  http.MethodGet,
			Handler: s.RedirectSvr.Redirect,
		},
	}
}
