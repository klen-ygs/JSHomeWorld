package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func addAImage(context *gin.Context) {
	file, err := context.FormFile("ImageData")
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     MODLEERR,
		})
		return
	}
	open, err := file.Open()
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     "",
		})
	}
	Image := DaoModle.Image{
		ImageData: make([]byte, file.Size),
	}
	_, err = open.Read(Image.ImageData)
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     "",
		})
	}
	Image.ShopId, _ = strconv.ParseUint(context.PostForm("shopId"), 10, 64)
	Image.Flag, _ = strconv.ParseBool(context.PostForm("flag"))
	err = DB.Create(&Image).Error
	if err != nil {

	}
}

func init() {

	Engine.POST("/addImage", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		addAImage(context)
	})
	accessOrigin("/addImage")
}
