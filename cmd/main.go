package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kangyueyue/short-url/cmd/app"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// main 主函数
func main() {
	appCmd := &cli.App{
		Name:        "short-url",
		Description: "短链接平台",
		Version:     "1.0.0",
		Commands:    nil,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Value: 9080,
				Usage: "端口",
			},
		},
		Action: run,
	}
	err := appCmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

// run 启动
func run(c *cli.Context) error {
	app, err := app.InitApp("./config/config.conf")
	if err != nil {
		panic(err)
	}
	port := c.Int("port")
	logrus.Infof("启动服务，监听端口：%d", port)
	if err := app.Serve(port); err != nil {
		panic(err)
	}
	defer app.Close()

	// 阻塞，保持长期运行
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) // 监听ctrl+c和kill命令

	// 阻塞等待
	<-sigCh
	logrus.Infof("接收到终止信号，开始关闭......")
	return nil
}
