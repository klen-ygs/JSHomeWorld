package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

	Engine.GET("/getImages", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		type images struct {
			ShopId uint64
			Id     int
		}
		request := images{}
		err := context.BindQuery(&request)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		if request.Id == 0 {
			count := 0
			DB.Model(&DaoModle.Image{}).Where("shop_id = ? and flag = false", request.ShopId).Count(&count)
			context.JSON(http.StatusOK, gin.H{
				"count": count,
			})
			return
		}
		image := DaoModle.Image{}
		err = DB.Where("shop_id = ? and flag = false", request.ShopId).Limit(request.Id).Offset(request.Id - 1).Find(&image).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.Data(http.StatusOK, "image/*", image.ImageData)
	})

	Engine.GET("/getImage", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		image := DaoModle.Image{}
		err := context.BindQuery(&image)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}

		err = DB.Where("shop_id = ? and flag = true", image.ShopId).First(&image).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.Data(http.StatusOK, "image/*", image.ImageData)
	})

	accessOrigin("/getImage")
}
