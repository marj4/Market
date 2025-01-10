package server

import (
	"Market/backend"
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
	router.LoadHTMLGlob("frontend/pages/*")

	router.GET("/ping", PingPage)

	router.GET("/", func(c *gin.Context) {
		data, err := db.GetAllProduct(DB)

		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Can`t reсieve data from database",
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Products": data,
		})
	})

	router.GET("register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	router.POST("/register", func(c *gin.Context) {

		login := c.DefaultPostForm("login", "")
		password := c.DefaultPostForm("password", "")
		email := c.DefaultPostForm("email", "")

		user := backend.User{
			Login:    login,
			Password: password,
			Email:    email,
		}

		if err := db.AddUser(DB, user); err != nil {
			c.JSON(500, gin.H{
				"Error": "Cant add user to db",
			})
			return
		}
	})

	return router
}

// Continue from here
func hash(password string) string {

}
