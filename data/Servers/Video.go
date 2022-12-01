package Servers

import (
	"data/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func getVideo(context *gin.Context) {
	path := context.Query("path")
	file, err := os.ReadFile(path)
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     VIDEOPATHERR,
		})
		return
	}
	context.Data(200, "video/mp4", file)
}

func init() {
	Engine.GET("/video/:opt", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		param := context.Param("opt")
		if param == "findVideo" {
			video := DaoModle.ShopVideo{}
			err := context.ShouldBindQuery(&video)
			if err != nil {
				context.JSON(http.StatusOK, response{
					Request: false,
					Err:     MODLEERR,
				})
				return
			}
			video.Path = "images\\" + strconv.FormatUint(video.ShopId, 10) + "\\video.mp4"
			_, err = os.Open(video.Path)
			if err != nil {
				context.JSON(http.StatusOK, response{
					Request: false,
					Err:     NOTFIND,
				})
				return
			}
			context.JSON(http.StatusOK, &video)
			return
		} else if param == "getVideo" {
			getVideo(context)
			return
		}
	})

	accessOrigin("/video")
}
