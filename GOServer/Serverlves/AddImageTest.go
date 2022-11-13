package Serverlves

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	Engine.GET("/addImageTest", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		cookie, _ := context.Cookie("login")
		fmt.Println("cookie: ", cookie)
	})
	Engine.POST("/addImageTest", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		file, err := context.FormFile("ImageData")
		if err != nil {
			return
		}
		open, err := file.Open()
		if err != nil {

		}
		bytes := make([]byte, file.Size)
		open.Read(bytes)
		context.Data(200, "image/png", bytes)
	})

	accessOrigin("/addImageTest")
}
