package main

import (
	"log"
	"os"
	c "decode-decrypt-playground/file/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}	

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	

	e.GET("/records", c.NewRecordController().GetAllRecords)
	e.GET("/records/:no", c.NewRecordController().GetRecordByNo)

	e.Start(":" + os.Getenv("PORT"))
}