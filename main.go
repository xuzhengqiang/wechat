package main

import (
	"os"
	_ "wechat/config"
	_ "wechat/database"
	"wechat/routes"
)

func main()  {
	r := routes.InitRouter()
	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port) // 监听并在 0.0.0.0:8080 上启动服务
}
