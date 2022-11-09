package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	fmt.Println("this is a update windows")
	route := gin.Default()

	route.POST("/upload", func(c *gin.Context) {

	})
}
