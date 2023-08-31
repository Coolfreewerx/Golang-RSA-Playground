package model

type WebAmloHeaders struct {
	Authorization string `json:"Authorization"` // case 1 "Basic <base64 encoded username:password>"
	Username      string `json:"Username"`	    // case 2 
	Password      string `json:"Password"`
	X_API_Key     string `json:"X-API-Key"`
	Host          string `json:"Host"`
}

type WebAmloResponse struct {
	// ReturnStatus  string `json:"returnStatus"`
	// ReturnMessage string `json:"returnMessage"`
	// KeyID         string `json:"keyId"`
	// FileName      string `json:"fileName"`
	// mimeType	  string `json:"mimeType"`
	// TotalRecord   int    `json:"totalRecord"`
	Result		  string `json:"Result"`
}