package middlewares

import "github.com/gin-contrib/cors"

// GetCORSConfig 返回CORS配置
func GetCORSConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许所有域
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	return config
}
