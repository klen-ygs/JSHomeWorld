package Servers

import (
	"data/DaoModle"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var logierMap sync.Map

func init() {
	var logier []DaoModle.User
	file, err := os.ReadFile("./user.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &logier)
	if err != nil {
		panic(err)
	}
	for i, _ := range logier {
		logierMap.Store(logier[i].Phone, &logier[i])
	}

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
		value, ok := logierMap.Load(User.Phone)
		if !ok {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     PHONEERR,
			})
			return
		}
		if User.Password != value.(*DaoModle.User).Password {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     PASSWORDERR,
			})
			return
		}
		context.SetCookie("login", strconv.FormatUint(loginCount, 10), 100000, "/", StaticURL, true, false)
		phoneMap.Store(loginCount, User.Phone)
		countMap.Store(User.Phone, loginCount)
		loginCount++
		context.JSON(http.StatusOK, response{
			Request: true,
			Err:     NONE,
		})
	})
	accessOrigin("/login")

	Engine.POST("/register", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		var User DaoModle.User
		err := context.BindJSON(&User)
		if err != nil {
			fmt.Println(err)
		}
		logierMap.Store(User.Phone, &User)
		context.JSON(http.StatusOK, gin.H{
			"register": true,
		})

	})
	accessOrigin("/register")

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
		value, ok := logierMap.Load(user.Phone)
		if !ok {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, value.(*DaoModle.User))
	})
	accessOrigin("/getUserInfo")
}
