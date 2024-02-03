package main

import (
	"fmt"
	"kubez_project/config"
	_ "kubez_project/controllers"
	"kubez_project/middlewares"
	"kubez_project/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(config.TimestampFormat)
	r := gin.Default()
	// r.Use(middlewares.JwtCheck)

	// 使用CORS中间件来配置跨域选项
	corsConfig := middlewares.GetCORSConfig()
	r.Use(cors.New(corsConfig))

	routers.RegisterRouters(r)
	r.Run()
}
