//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"github.com/kangyueyue/road"

	"github.com/kangyueyue/short-url/internal/config"
	"github.com/kangyueyue/short-url/internal/httpserver"
)

var AppSet = wire.NewSet(
	NewApp,
	road.InitRoad,
	config.NewConfig,
	httpserver.InitialHttpServer,
)

var DBComponentSet = wire.NewSet(
	config.ProvideStore,
	config.ProvideRedisClient,
)

// InitApp 初始化应用
func InitApp(road string) (*App, error) {
	wire.Build(
		AppSet,
		DBComponentSet,
	)
	return &App{}, nil // 没有实际意义，只是为了不报错
}
