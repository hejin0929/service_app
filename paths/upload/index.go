package upload

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func Upload(c *gin.Context) {

	_ = c.Request.ParseMultipartForm(32 << 20)

	file, err := c.FormFile("name")

	if err != nil {
		c.JSON(http.StatusInternalServerError, "文件读取错误")
	}

	createFile, _ := os.OpenFile("./assets/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	defer createFile.Close()

	data, _ := file.Open()

	io.Copy(createFile, data)

	res := struct {
		Url string `json:"url"`
	}{
		Url: "http://localhost:8080/api/assets/" + file.Filename,
	}

	c.JSON(http.StatusOK, res)
}
