package main

import (
	"log"
)

func main() {
	// 加载配置
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}
	// 初始化服务、注入依赖等
	server, err := initServer()
	if err != nil {
		log.Fatal(err)
	}
	// 服务启动
	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
