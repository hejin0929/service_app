package model

type HTTPError struct {
	Error   string `json:"error" binding:"required"`
	Message string `json:"message" binding:"required"`
}
