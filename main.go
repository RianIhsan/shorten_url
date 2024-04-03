package main

import (
	"fmt"

	"github.com/RianIhsan/shorten_url/config"
	"github.com/RianIhsan/shorten_url/feature/url/handler"
	"github.com/RianIhsan/shorten_url/feature/url/repository"
	"github.com/RianIhsan/shorten_url/feature/url/service"
	"github.com/RianIhsan/shorten_url/helper/cache/redis"
	"github.com/RianIhsan/shorten_url/helper/database"
	"github.com/RianIhsan/shorten_url/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	var loadConfig = config.RunConfig()
	db := database.NewConnectionDB(*loadConfig)
	rdb := redis.NewRedisClient(*loadConfig)
	database.Migrate(db)
	urlRepo := repository.NewURLRepository(db)
	urlSvc := service.NewURLService(urlRepo, rdb)
	urlHandler := handler.NewURLHandler(urlSvc)

	app.Use(corsMiddleware())

	routes.URlRoute(app, urlHandler)
	addr := fmt.Sprintf(":%d", loadConfig.AppPort)
	app.Run(addr)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
