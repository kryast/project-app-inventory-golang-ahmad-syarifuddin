package model

type PaginatedResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"` // Pastikan field ini ada
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalItems int         `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}
