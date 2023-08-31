package main

import (
	c "decode-decrypt-playground/controller"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"

	_ "decode-decrypt-playground/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API for Golang-RSAPlayground
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @host localhost:1122
// @BasePath /
// @schemes http
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Encrypt/Decrypt file with RSA and encode/decode encrypted file with base64
	// ciphers.CipherFileWithRSAandBase64(os.Getenv("TAM_PUBLIC_KEY_PATH"), os.Getenv("TAM_PRIVATE_KEY_PATH"), 
	// 								   os.Getenv("FILE_ZIP_PATH"), os.Getenv("TAM_PASSPHRASE"), os.Getenv("ENCRYPT_FILE_PATH"), os.Getenv("DECRYPT_FILE_PATH"))

	e := echo.New()

	// Middleware for adding custom headers
	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// WEB AMLO
	e.POST("/web-amlo", c.NewWebAmloController().WebAmloRequest)
	
	// DATABASE
	e.GET("/records", c.NewRecordController().GetAllRecords)
	e.GET("/records/:no", c.NewRecordController().GetRecordByNo)

	e.Start(":" + os.Getenv("PORT"))
}
