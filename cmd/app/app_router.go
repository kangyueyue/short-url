package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/cmd/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// registerRouter 注册路由
func (a *App) registerRouter() {
	g := a.engine
	for path, c := range a.httpServer.GetRounter() {
		g.Handle(c.Method, path, c.Handler)
	}
}

// registerSwagger 注册swagger
func (a *App) registerSwagger(r gin.IRouter) {
	// API文档访问地址: http://host/swagger/index.html
	// 注解定义可参考 https://github.com/swaggo/swag#declarative-comments-format
	// 样例 https://github.com/swaggo/swag/blob/master/example/basic/api/api.go
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Title = "Swagger接口文档"
	docs.SwaggerInfo.Description = "接口文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
