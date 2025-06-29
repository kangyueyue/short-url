package app

// registerRouter 注册路由
func (a *App) registerRouter() {
	g := a.engine
	for path, c := range a.httpServer.GetRounter() {
		g.Handle(c.Method, path, c.Handler)
	}
}
