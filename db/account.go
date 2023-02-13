package db

import "gorm.io/gorm"

// User 用户信息库
type User struct {
	gorm.Model
	Uuid       string `json:"uuid"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Sex        int    `json:"sex"`
	Introduce  string `json:"introduce"`
	CodeNumber string `json:"code_number"`
}

// Code 验证码库
type Code struct {
	gorm.Model
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
