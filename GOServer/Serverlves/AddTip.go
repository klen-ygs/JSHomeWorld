package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Engine.POST("/addTip", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		request := DaoModle.ShopList{}
		err := context.BindJSON(&request)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		getPhone(context, &request)
		err = DB.Model(&request).Where("phone = ?", request.Phone).Update(&request).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}

	})

	Engine.GET("/addTip", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		request := DaoModle.ShopList{}
		err := context.BindQuery(&request)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		getPhone(context, &request)
		err = DB.Find(&request).Error
		if err != nil {

			return
		}
		context.JSON(http.StatusOK, &request)
	})
	accessOrigin("/addTip")
}
