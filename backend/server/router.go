package server

import (
	"Market/backend"
	"Market/backend/db"
	"Market/config"
	error2 "Market/error"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanorand/nanorand"
	cors "github.com/rs/cors/wrapper/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/smtp"
	"regexp"
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

	router.GET("/register", func(c *gin.Context) { c.HTML(http.StatusOK, "register.html", nil) })
	router.POST("/register", func(c *gin.Context) {

		//Получаю данные с формы
		login := c.PostForm("login")
		password := c.PostForm("password")
		email := c.PostForm("email")

		//Дополнительная проверка данных на стороне сервера
		if err := validateUserData(login, email, password); err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": err.Error()})
			return
		}

		//Получаю логин и почту из БД, для того чтобы проверить, существует ли пользователь с введеными данными
		loginsEmail, err := db.GetAllLoginAndEmail(DB)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": "Can`t receive users from db"})
			return
		}

		//Проверяем
		for _, log := range loginsEmail {
			if log.Login == login && log.Email == email {
				c.HTML(http.StatusOK, "register.html", gin.H{"Error": "User with this data already exists"})
			} else if log.Login == login {
				c.HTML(http.StatusOK, "register.html", gin.H{"Error": "User with this login already exists"})
			} else if log.Email == email {
				c.HTML(http.StatusOK, "register.html", gin.H{"Error": "User with this email already exists"})
			}
		}

		//Хеширую введённый пароль
		_, hashPassword, err := hash(password)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": err.Error()})
			return
		}

		user := backend.User{
			Login:    login,
			Password: hashPassword,
			Email:    email,
		}

		//Добавляю пользователя в БД
		if err := db.AddUser(DB, user); err != nil {
			log.Fatal(err)

			c.JSON(500, gin.H{
				"Error": "Cant add user to db",
			})
			return
		}

		//Перекидываю пользователя на главную страницу
		c.Redirect(http.StatusSeeOther, "/")

	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", func(c *gin.Context) {
		login := c.PostForm("login")
		password := c.PostForm("password")

		//Hash password user from sing in
		hash, _, err := hash(password)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": "Error for hash..//"})
			return
		}

		//Get hash-password from DB for check
		PasswordUser, err := db.GetUser(DB, login)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"Error": "Error for database"})
			return
		}

		//Check password user with password from DB
		if err := bcrypt.CompareHashAndPassword([]byte(PasswordUser), hash); err != nil {
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

func GenerateCodeForEmail() string {
	code, _ := nanorand.Gen(6)
	return code
}

func SendCodeToEmail(email string, code string) error {
	//Загружаю конфиг, из которого буду брать почту, от которой будут рассылаться коды для двухфакторки
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//Готовлю письмо
	subject := "Email Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", GenerateCodeForEmail())
	message := []byte("Subject: " + subject + "\r\n\r\n" + body)

	//Параметры SMTP-сервера
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Ввожу данные для авторизации на почте, от которой будут рассылаться коды
	auth := smtp.PlainAuth("", cfg.Email, cfg.APP_PASSWORD, smtpHost)

	//Отправляю код
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, cfg.Email, []string{email}, message); err != nil {
		log.Fatal(err)

	}

	return nil

}

func validateUserData(login, email, password string) error {
	// Проверка логина
	if len(login) < 3 || len(login) > 16 {
		return errors.New("login must be between 3 and 16 characters")
	}
	loginRegex := `^[a-zA-Z0-9._-]+$`
	if !regexp.MustCompile(loginRegex).MatchString(login) {
		return errors.New("login can only contain letters, numbers, '.', '_', and '-'")
	}

	// Проверка пароля
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must include at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must include at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must include at least one number")
	}
	if !regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password) {
		return errors.New("password must include at least one special character")
	}

	return nil
}
