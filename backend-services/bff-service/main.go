package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	IsAuthorization bool `json:"isAuthorization"`
}

type UserContext struct {
	User User `json:"user"`
}

var mockUserContext = UserContext{
	User: User{
		Username: "Alisher Sharipov",
		IsAuthorization: true,
	},
} 

func getContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"context": mockUserContext,
	})
}

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/bff/context", getContext)

	r.Run(":8081")
}
