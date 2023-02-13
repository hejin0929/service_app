package model

type Account struct {
	Phone string `json:"phone"`
	Uuid  string `json:"uuid"`
}

type AccountResult struct {
	Uuid  string `json:"uuid"`
	Phone string `json:"phone"`
}

type AccountRequest struct {
	Password string `json:"password"`
	Code     string `json:"code"`
}

type LoginResult struct {
	Token string `json:"token" binding:"required"`
	Uuid  string `json:"uuid" binding:"required"`
	Time  int    `json:"time" binding:"required"`
}
