package controller

import (
	"crypto/tls"
	m "decode-decrypt-playground/model"
	"decode-decrypt-playground/ciphers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type WebAmloController struct {}

func NewWebAmloController() *WebAmloController {
	return &WebAmloController{}
} 

func (c *WebAmloController) WebAmloRequest (context echo.Context) error {

	url := "/v1/hr-02/1.1"

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	headers := m.WebAmloHeaders{
		Authorization: os.Getenv("AUTHORIZATION"),
		X_API_Key: os.Getenv("X_API_KEY"),
		Host: os.Getenv("HOST"),
	}
	
	// read file .pfx of certificate
	pfxData, err := ioutil.ReadFile(os.Getenv("PFX_PATH"))
	if err != nil {
		fmt.Println("Error reading .pfx file", err)
	}

	// check passphase for .pfx file
	pfx, err := tls.X509KeyPair(pfxData, []byte(os.Getenv("PFX_PASSPHASE")))
	if err != nil {
		fmt.Println("Error parsing .pfx file", err)
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// make connect by use resty 
	client := resty.New()
	client.SetCertificates(pfx)
	client.SetDebug(true)

	// set header for request
	resp, err := client.R().
		SetHeaders(map[string]string{	
			"X-API-Key": headers.X_API_Key,
			"Host": headers.Host,
			"Authorization": headers.Authorization, // case 1
		}).
		SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD")). // case 2
		Post(headers.Host + url)
		// Post(url) 

	if err != nil {
		fmt.Println("Error connecting to server", err)
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	defer resp.RawResponse.Body.Close()
	
	// return context.JSON(http.StatusOK, resp)

	// read response body
	// resp_body, err := ioutil.ReadAll(resp.RawResponse.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body", err)
	// 	return context.JSON(http.StatusInternalServerError, err.Error()) 
	// }

	// map json response body "result" to struct WebAmloResponse
	response := m.WebAmloResponse{}
	err = json.Unmarshal(resp.Body(), &response)

	ciphers.DecodeDecryptFile(response.Result, os.Getenv("PRIVATE_KEY_DATA"), os.Getenv("PASSPHRASE"), os.Getenv("ENCRYPT_FILE_PATH"), os.Getenv("DECRYPT_FILE_PATH"))

	return context.JSON(http.StatusOK, "Success request to Web AMLO")
}
