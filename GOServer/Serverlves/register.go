package Serverlves

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 用户模型
type user struct {
	Phone    string
	Name     string
	Password string
}

func init() {
	Engine.POST("/register", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		var P user
		err := context.BindJSON(&P)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", P)
	})
	Engine.OPTIONS("/register", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,XFILENAME,XFILECATEGORY,XFILESIZE")
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	})
}
