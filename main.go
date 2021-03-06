package main

import (
	"fmt"

	"giligili-api/conf"
	"giligili-api/server"
)

func main() {
	fmt.Println("go giligili")
	// 从配置文件读取配置
	conf.Init()

	//装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
