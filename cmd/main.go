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
}
