package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {

	Engine.GET("/getUserInfo", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		cookie, err := context.Cookie("login")
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		count, _ := strconv.ParseUint(cookie, 10, 64)
		phoneLoad, ok := phoneMap.Load(count)
		if !ok {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     "CookieFail",
			})
			return
		}
		user := DaoModle.User{
			Phone:    phoneLoad.(string),
			Name:     "",
			Password: "",
		}
		err = DB.Model(&user).Select("name").Where("phone = ?", user.Phone).Find(&user).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, &user)
	})
	accessOrigin("/getUserInfo")
}
