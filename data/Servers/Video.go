package Servers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Engine.GET("/video/:opt", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		param := context.Param("opt")
		if param == "findVideo" {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		} else if param == "getVideo" {
			//getVideo(context) //假数据后台暂不支持视频播放
			return
		}
	})

	accessOrigin("/video")
}
