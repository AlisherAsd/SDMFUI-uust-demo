package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Item struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description  string `json:"description"`
	Price int `json:"price"`
}


type ItemCreate struct {
	Title string `json:"title"`
	Description  string `json:"description"`
	Price int `json:"price"`
}


func getItemsList(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := p.Query(c.Request.Context(), "SELECT * FROM items")

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		items, err := pgx.CollectRows(rows, pgx.RowToStructByName[Item])

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	}

}


func addItem(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item ItemCreate

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := p.Exec(
			c.Request.Context(),
			"INSERT INTO items (title, description, price) VALUES ($1, $2, $3)",
			item.Title, item.Description, item.Price,
		)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
		})
	}
}


type Config struct {
	DatabaseURL string
}

func Load() *Config {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@layout-service-postgres:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	return &Config{DatabaseURL: dsn}
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

		cfg := Load()

	pool, err :=  pgxpool.New(context.Background(), cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer pool.Close()


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/items-service/list", getItemsList(pool))

	r.POST("/items-service/item", addItem(pool))

	r.Run(":8082")
}
