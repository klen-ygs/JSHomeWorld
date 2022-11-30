package Serverlves

import (
	"GOServer/DaoModle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Engine.POST("/register", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		var User DaoModle.User
		err := context.BindJSON(&User)
		if err != nil {
			fmt.Println(err)
		}
		err = DB.Create(&User).Error
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"register": false,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"register": true,
			})
		}
	})
	accessOrigin("/register")
}
