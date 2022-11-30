package Serverlves

import (
	"GOServer/DaoModle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 注册 /login
// 提供前端登录接口
func init() {
	Engine.POST("/login", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		context.SetSameSite(http.SameSiteNoneMode)
		var User DaoModle.User
		err := context.BindJSON(&User)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		context.SetCookie("login", strconv.FormatUint(loginCount, 10), 100000, "/", StaticURL, true, false)
		phoneMap.Store(loginCount, User.Phone)
		countMap.Store(User.Phone, loginCount)
		loginCount++
		err = DB.Model(&User).Where("phone = ? and password = ?", User.Phone, User.Password).First(&User).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			fmt.Println(err)
			return
		}
		context.JSON(http.StatusOK, response{
			Request: true,
			Err:     NONE,
		})
	})
	accessOrigin("/login")
	Engine.DELETE("/login", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "DELETE")
		user := DaoModle.User{}
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		count, ok := countMap.Load(user.Phone)
		if !ok {
			return
		}
		countMap.Delete(user.Phone)
		phoneMap.Delete(count)
		context.SetCookie("login", "", -1, "/", StaticURL, true, false)
	})

}
