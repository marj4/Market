package server

import (
	"Market/backend"
	"Market/backend/db"
	error2 "Market/error"
	"database/sql"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
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
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": "Can`t reсieve data from database"})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{"Products": data})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	router.POST("/register", func(c *gin.Context) {

		login := c.PostForm("login")
		password := c.PostForm("password")
		email := c.PostForm("email")

		_, hashPassword, err := hash(password)
		if err != nil {
			log.Fatal(err)

			c.JSON(500, gin.H{
				"Error": err.Error(),
			})
		}

		user := backend.User{
			Login:    login,
			Password: hashPassword,
			Email:    email,
		}

		if err := db.AddUser(DB, user); err != nil {
			log.Fatal(err)

			c.JSON(500, gin.H{
				"Error": "Cant add user to db",
			})
			return
		}

		c.Redirect(http.StatusSeeOther, "/")

	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", func(c *gin.Context) {
		login := c.PostForm("login")
		password := c.PostForm("password")

		//Get hash-password from DB for check
		UserPassword, err := db.GetUser(DB, login)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": "Error for database"})
			return
		}

		//Hash password user
		hash, _, err := hash(password)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": "Error for hash..//"})
			return
		}

		//Check password user with password from DB
		if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
			log.Fatal(err)
			c.JSON(401, gin.H{"Error": "Invalid username or password"})
			return
		}

		//token,err := GenerateToken()

		c.Redirect(http.StatusSeeOther, "/")

	})

	return router
}

// Continue from here
func hash(password string) ([]byte, string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, "", error2.Wrap("Cant hash password", err)
	}
	return hash, string(hash), nil
}

//func checkPassword()
