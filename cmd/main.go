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

	server, err := initServer()
	if err != nil {
		log.Fatal(err)
	}

	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
