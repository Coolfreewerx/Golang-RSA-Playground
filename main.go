package main

import (
	// "decode-decrypt-playground/ciphers"

	"log"
	"os"
	c "decode-decrypt-playground/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger" 
	_ "decode-decrypt-playground/docs"


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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	
	e.GET("/records", c.NewRecordController().GetAllRecords)
	e.GET("/records/:no", c.NewRecordController().GetRecordByNo)

	e.Start(":" + os.Getenv("PORT"))
}
