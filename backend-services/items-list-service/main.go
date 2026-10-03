package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description  string `json:"description"`
	Price int `json:"price"`
}



var mockItemsList = []Item{
	Item{
		Id: 1,
		Title: "item 1",
		Description: "description 1",
		Price: 1000,
	},
	Item{
		Id: 2,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 3,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 4,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 5,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 6,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 7,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 8,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 9,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 10,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
	Item{
		Id: 11,
		Title: "item 2",
		Description: "description 2",
		Price: 2000,
	},
} 

func getItemsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": mockItemsList,
	})
}

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
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

	r.GET("/items-service/list", getItemsList)

	r.Run(":8082")
}
