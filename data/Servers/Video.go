package Servers

import (
	"data/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getVideo(context *gin.Context) {
	RangeStr := context.GetHeader("Range")
	if len(RangeStr) == 1 { // 用普通Get请求
		RangeStr = "bytes=0-"
	}
	RangeNum := strings.Split(RangeStr, "=")[1]
	RangeIdx := strings.Split(RangeNum, "-")[0]
	seek, _ := strconv.ParseInt(RangeIdx, 10, 64)
	path := context.Query("path")
	file, err := os.ReadFile(path)
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     VIDEOPATHERR,
		})
		return
	}
	rep := "bytes "
	rep += strconv.FormatInt(seek, 10) + "-" + strconv.FormatInt(int64(len(file)-1), 10) + "/" + strconv.FormatInt(int64(len(file)), 10)
	context.Writer.Header().Add("Content-Range", rep)
	context.Data(206, "video/mp4", file[seek:])
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
