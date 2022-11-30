package Servers

import (
	"data/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var shopListMap sync.Map

func init() {
	Engine.POST("/addTip", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		request := DaoModle.ShopList{}
		err := context.BindJSON(&request)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		getPhone(context, &request)
		shopListMap.Store(request.Phone, &request)
		context.JSON(http.StatusOK, response{
			Request: true,
			Err:     NONE,
		})
	})

	Engine.GET("/addTip", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		request := DaoModle.ShopList{
			Phone: "",
			List:  "[]",
		}
		err := context.BindQuery(&request)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		getPhone(context, &request)
		value, ok := shopListMap.Load(request.Phone)
		if ok {
			request = *value.(*DaoModle.ShopList)
		}
		context.JSON(http.StatusOK, &request)
	})
	accessOrigin("/addTip")
}
