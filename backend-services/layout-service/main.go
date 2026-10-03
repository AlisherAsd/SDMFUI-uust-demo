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

type Widget struct {
	MfName        string `json:"mfName"`
	ComponentName string `json:"componentName"`
	EntryUrl      string `json:"entryUrl"`
	Id            int    `json:"id"`
}

type LayoutWidget struct {
	Widget
	Level float64 `json:"level"`
}

type Page struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreatePage struct {
	Name string `json:"name"`
}

type WidgetPage struct {
	PageId   int     `json:"pageId"`
	WidgetId int     `json:"widgetId"`
	Level    float64 `json:"level"`
}

func getWidgets(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := p.Query(c.Request.Context(), "SELECT * FROM widgets")

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		widgets, err := pgx.CollectRows(rows, pgx.RowToStructByName[Widget])

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"widgets": widgets,
		})
	}
}

func createWidget(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var widget Widget

		if err := c.ShouldBindJSON(&widget); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := p.Exec(
			c.Request.Context(),
			`INSERT INTO widgets (mf_name, component_name, entry_url) VALUES ($1, $2, $3)`,
			widget.MfName, widget.ComponentName, widget.EntryUrl,
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
		return
	}
}

func getPages(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := p.Query(c.Request.Context(), "SELECT * FROM pages")

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		pages, err := pgx.CollectRows(rows, pgx.RowToStructByName[Page])

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"pages": pages,
		})
	}
}

func createPage(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page CreatePage

		if err := c.ShouldBindJSON(&page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := p.Exec(
			c.Request.Context(),
			"INSERT INTO pages (name) VALUES ($1)",
			page.Name,
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
		return
	}
}

func getLayout(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := p.Query(
			c.Request.Context(),
			"SELECT w.id, pw.level, w.mf_name, w.component_name, w.entry_url FROM widgets w JOIN pages_widgets pw ON pw.widget_id = w.id JOIN pages p ON pw.page_id = p.id WHERE p.name = $1",
			c.Param("pageName"),
		)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		widgets, err := pgx.CollectRows(rows, pgx.RowToStructByName[LayoutWidget])

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"widgets": widgets,
		})
	}
}

func addWidgetOnPage(p *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wp WidgetPage

		if err := c.ShouldBindJSON(&wp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("INSERT pages_widgets: pageId=%d widgetId=%d level=%d",
			wp.PageId, wp.WidgetId, wp.Level)
		_, err := p.Exec(
			c.Request.Context(),
			`INSERT INTO pages_widgets (page_id, widget_id, level) VALUES ($1, $2, $3)`,
			wp.PageId, wp.WidgetId, wp.Level,
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

	r.GET("/widgets", getWidgets(pool))

	r.POST("/widgets", createWidget(pool))

	r.GET("/widgets/:pageName", getLayout(pool))

	r.GET("/pages", getPages(pool))

	r.POST("/pages", createPage(pool))

	r.POST("/pages/widget", addWidgetOnPage(pool))

	r.Run()

}
