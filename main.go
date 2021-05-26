package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Client godoc
type Client chan string

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	router.GET("/helloworld", func(c *gin.Context) {
		cmd := redisClient.Get("cache-msg")
		if cmd.Err() != nil {
			c.String(http.StatusInternalServerError, cmd.Err().Error())
		}
		c.String(http.StatusOK, cmd.Val())
	})
	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// For Testing server is ready
	time.Sleep(time.Second * 3)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run(":" + port)
}
