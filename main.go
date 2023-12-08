package main

import (
	"kubez_project/middlewares"
	"kubez_project/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(middlewares.JwtCheck)
	routers.RegisterRouters(r)

	r.Run()
}
