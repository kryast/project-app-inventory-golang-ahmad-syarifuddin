package model

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       Item   `json:"data"`
}
