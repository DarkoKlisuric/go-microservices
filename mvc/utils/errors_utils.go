package utils

type AppError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}
