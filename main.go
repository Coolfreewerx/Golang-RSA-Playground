package main

import (
	"decode-decrypt-playground/ciphers"

	"log"
	"os"
	c "decode-decrypt-playground/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	
	// Encrypt/Decrypt message with RSA and encode/decode encrypted message with base64
	ciphers.CipherWithRSAandBase64(os.Getenv("TAM_PUBLIC_KEY"), os.Getenv("TAM_PRIVATE_KEY"), os.Getenv("MESSAGE"), os.Getenv("TAM_PASSPHRASE"))

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	
	e.GET("/records", c.NewRecordController().GetAllRecords)
	e.GET("/records/:no", c.NewRecordController().GetRecordByNo)

	e.Start(":" + os.Getenv("PORT"))
}
