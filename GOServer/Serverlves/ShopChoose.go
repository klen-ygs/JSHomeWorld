package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

	Engine.GET("/shopChoose", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", AJAXOrigin)
		choose := DaoModle.ShopChoose{}
		err := context.ShouldBindQuery(&choose)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		err = DB.Where("shop_id = ?", choose.ShopId).Find(&choose).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, &choose)
	})

	Engine.POST("/shopChoose", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", AJAXOrigin)
		choose := DaoModle.ShopChoose{}
		err := context.ShouldBindQuery(&choose)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		err = DB.Create(&choose).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, response{
			Request: true,
			Err:     NONE,
		})
	})

	accessOrigin("/shopChoose")
}
