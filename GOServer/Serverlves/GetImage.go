package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

	Engine.GET("/getImage", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		image := DaoModle.Image{}
		err := context.BindJSON(&image)
		if err != nil {
			context.JSON(http.StatusOK, response{
				request: false,
				Err:     MODLEERR,
			})
			return
		}
		err = DB.Where("id = ?", image.Id).First(&image).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, &image)
	})

	accessOrigin("/getImage")
}
