// package main

// import (
// 	"fmt"
// 	"kubez_project/config"
// 	"kubez_project/middlewares"
// 	"kubez_project/routers"
// 	"kubez_project/utils/jwts"
// 	"kubez_project/utils/logs"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	logs.Info(nil, "test")
// 	fmt.Println(config.JwtSecret)

// 	token, _ := jwts.GenToken("admin")
// 	clam, err := jwts.ParseToken(token)
// 	fmt.Println(clam, err)

// 	r := gin.Default()
// 	routers.RegisterRouters(r)
// 	r.Use(middlewares.JwtCheck)
// 	r.Run()
// }
