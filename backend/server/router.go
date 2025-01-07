package server

import (
	"Market/backend/db"
	"database/sql"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
)

func LoadRouter(DB *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:63342"}, // Замени на адрес фронтенда
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Указываем путь к шаблонам
	router.LoadHTMLGlob("frontend/index.html")

	router.Static("/static/styles.css", "./frontend")

	router.GET("/ping", PingPage)

	router.GET("/shop.ru", func(c *gin.Context) {
		data, err := db.GetAllProduct(DB) // Здесь ID = 1, ты можешь менять его динамически

		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Can`t retrieve data from database",
			})
			return
		}

		// Передаём HTML с динамическими данными
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Products": data,
		})
	})

	return router
}
