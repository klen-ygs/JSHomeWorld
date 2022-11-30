package Servers

import (
	"data/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func init() {
	Engine.GET("/getImages", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
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
			count := 3
			context.JSON(http.StatusOK, gin.H{
				"count": count,
			})
			return
		}
		image := DaoModle.Image{}
		path := "./images/" + strconv.FormatUint(request.ShopId, 10) + "/" + strconv.FormatInt(int64(request.Id), 10) + ".png"
		file, err := os.ReadFile(path)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		image.ImageData = file
		context.Data(http.StatusOK, "image/*", image.ImageData)
	})
	accessOrigin("/getImages")

	Engine.GET("/getImage", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		image := DaoModle.Image{}
		err := context.BindQuery(&image)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		path := "./images/" + strconv.FormatUint(image.ShopId, 10) + "/0.png"
		file, err := os.ReadFile(path)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		image.ImageData = file
		context.Data(http.StatusOK, "image/*", image.ImageData)
	})

	accessOrigin("/getImage")

}
