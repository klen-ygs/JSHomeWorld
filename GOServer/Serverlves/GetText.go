package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Engine.GET("/getText", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		shopText := DaoModle.Text{}
		err := context.BindJSON(&shopText)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		err = DB.Where("id = ?", shopText.Id).First(&shopText).Error
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, &shopText)
	})
	accessOrigin("/getText")
}
