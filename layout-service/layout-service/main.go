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
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func Load() *Config {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	return &Config{DatabaseURL: dsn}
}

type Remote struct {
	MfName        string `json:"mfName"`
	ComponentName string `json:"componentName"`
	EntryUrl      string `json:"entryUrl"`
}

type CreateRemote struct {
	Remote
	Page string
}

func getRemotes(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := p.Query(
			c.Request.Context(),
			"SELECT mf_name, component_name, entry_url FROM remotes where page = $1",
			c.Param("name"),
		)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		remotes, err := pgx.CollectRows(rows, pgx.RowToStructByName[Remote])

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"remotes": remotes,
		})
	}
}

func createRemote(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var remote CreateRemote

		if err := c.ShouldBindJSON(&remote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := p.Exec(
			c.Request.Context(),
			`INSERT INTO remotes (page, mf_name, component_name, entry_url) VALUES ($1, $2, $3, $4)`,
			remote.Page, remote.MfName, remote.ComponentName, remote.EntryUrl,
		)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadGateway, gin.H{
			"success": true,
		})
		return
	}
}

func main() {
	_ = godotenv.Load()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cfg := Load()
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer pool.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/remotes/:name", getRemotes(pool))

	r.POST("/remotes", createRemote(pool))

	r.Run()

}
