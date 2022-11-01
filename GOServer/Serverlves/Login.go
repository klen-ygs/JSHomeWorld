package Serverlves

import (
	"GOServer/DaoModle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册 /login
// 提供前端登录接口
func init() {
	Engine.POST("/login", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		var User DaoModle.User
		err := context.BindJSON(&User)
		if err != nil {
			context.JSON(http.StatusOK, response{
				request: false,
				Err:     MODLEERR,
			})
			return
		}
		err = DB.Model(&User).Where("phone = ? and password = ?", User.Phone, User.Password).First(&User).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				request: false,
				Err:     NOTFIND,
			})
			fmt.Println(err)
			return
		}
		context.JSON(http.StatusOK, response{
			request: true,
			Err:     NONE,
		})
	})
	accessOrigin("/login")

}
