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

	router.Static("/static", "./frontend") // Все статические файлы будут доступны через /static

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "true",
		})
	})

	router.GET("/shop.ru", func(c *gin.Context) {
		data, err := db.GetAllProduct(DB, 3) // Здесь ID = 1, ты можешь менять его динамически

		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Can`t retrieve data from database",
			})
			return
		}

		// Передаём HTML с динамическими данными
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name":        data.Name,
			"description": data.Description,
			"price":       data.Price,
			"picture":     data.Picture_URL,
		})
	})

	return router
}
