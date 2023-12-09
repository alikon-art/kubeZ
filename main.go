package main

import (
	"fmt"
	"kubez_project/config"
	_ "kubez_project/controllers"
	// "kubez_project/middlewares"
	"kubez_project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(config.TimestampFormat)
	r := gin.Default()
	// r.Use(middlewares.JwtCheck)
	routers.RegisterRouters(r)
	r.Run()
}
