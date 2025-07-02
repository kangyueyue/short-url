//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package httpserver

import (
	"github.com/google/wire"
	"github.com/kangyueyue/short-url/internal/httpserver/hello"
	"github.com/kangyueyue/short-url/internal/httpserver/short_url"
	"github.com/kangyueyue/short-url/internal/infrastructure/store"
)

var SvrSet = wire.NewSet(
	hello.NewHelloSvr,
	short_url.NewShortUrlSvr,
)

// InitialHttpServer 初始化服务
func InitialHttpServer(store *store.Store) *Server {
	wire.Build(
		NewServer,
		SvrSet,
	)
	return &Server{}
}
