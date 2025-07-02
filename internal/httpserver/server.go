package httpserver

import (
	"github.com/kangyueyue/short-url/internal/httpserver/hello"
	"github.com/kangyueyue/short-url/internal/httpserver/short_url"
)

// Server 服务
type Server struct {
	*hello.HelloSvr
	*short_url.ShortUrlSvr
}

// NewServer 创建服务
func NewServer(h *hello.HelloSvr,
	su *short_url.ShortUrlSvr,
) *Server {
	return &Server{
		HelloSvr:    h,
		ShortUrlSvr: su,
	}
}
