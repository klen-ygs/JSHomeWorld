package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Engine.GET("/getShop", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		Shop := DaoModle.Shop{}
		err := context.BindJSON(&Shop)
		if err != nil {
			context.JSON(http.StatusOK, response{
				request: false,
				Err:     MODLEERR,
			})
			return
		}
		err = DB.Where("id = ?", Shop.Id).First(&Shop).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, &Shop)
	})
	accessOrigin("/getShop")

}
