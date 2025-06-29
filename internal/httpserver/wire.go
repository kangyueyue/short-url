//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package httpserver

import (
	"github.com/google/wire"
	"github.com/kangyueyue/short-url/internal/httpserver/hello"
)

var SvrSet = wire.NewSet(
	hello.NewHelloSvr,
)

// InitialHttpServer 初始化服务
func InitialHttpServer() *Server {
	wire.Build(
		NewServer,
		SvrSet,
	)
	return &Server{}
}
