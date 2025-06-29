package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/road"
	"github.com/kangyueyue/short-url/internal/config"
	"github.com/kangyueyue/short-url/internal/httpserver"
	"github.com/kangyueyue/short-url/internal/infrastructure/log"
	"github.com/kangyueyue/short-url/internal/infrastructure/store"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// App 应用
type App struct {
	road       *road.Road         // road
	cfg        *config.Config     // 配置类
	httpServer *httpserver.Server // http server
	engine     *gin.Engine        // gin engine
	// db
	store       *store.Store  // store
	redisClient *redis.Client // redis client
}

// NewApp 新建应用
func NewApp(
	road *road.Road,
	cfg *config.Config,
	httpServer *httpserver.Server,
	store *store.Store,
	redisClient *redis.Client,
) *App {
	return &App{
		road:        road,
		cfg:         cfg,
		httpServer:  httpServer,
		engine:      gin.New(),
		store:       store,
		redisClient: redisClient,
	}
}

// Serve 启动服务
func (a *App) Serve(port int) error {
	// 设置日志
	log.InitLog()
	a.registerRouter() // 注册路由
	svr := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           a.engine,
		ReadHeaderTimeout: 5 * time.Second,
	}
	// 异步启动http server
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			logrus.Fatalf("listen err: %v", err)
		}
	}()
	return nil
}

// Close 关闭
func (a *App) Close() {
}
