package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"service_app/db"
	"service_app/docs"
	account "service_app/paths/account"
	"service_app/paths/upload"
	"strings"
)

func init() {

}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.basic	BasicAuth
func main() {

	HttpApp := gin.Default()

	docs.SwaggerInfo.BasePath = "api"

	db.ConnectDB()

	HttpApp.Use(Cors())

	route := HttpApp.Group("/api")

	route.GET("/doc", func(c *gin.Context) {
		file, _ := os.ReadFile("./api.json")
		_, _ = c.Writer.Write(file)
	})

	route.GET("/assets/*any", func(c *gin.Context) {

		lists := strings.Split(c.Request.URL.Path, "/")

		file, err := os.ReadFile("./assets/" + lists[len(lists)-1])

		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}
		_, _ = c.Writer.Write(file)
	})

	login := HttpApp.Group("/account")

	// 注册接口
	login.POST("/:phone", account.RegisterPaths)

	route.POST("/upload", upload.Upload)

	HttpApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)

	HttpApp.Run(":8080").Error()
}

// Cors //// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") // 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")

			c.Set("Access-Control-Allow-Origin", "*")    //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")    //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "multipart/form-data") // 设置返回格式是json
			c.Set("content-type", "text/plain; charset=utf-8")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
