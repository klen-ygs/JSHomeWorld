package Servers

import (
	"data/DaoModle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

var chooseMap sync.Map

func init() {
	file, err := os.ReadFile("./shop_choose.json")
	if err != nil {
		panic(err)
	}
	var chooses []DaoModle.ShopChoose
	err = json.Unmarshal(file, &chooses)
	if err != nil {
		panic(err)
	}
	for i, _ := range chooses {
		chooseMap.Store(chooses[i].ShopId, &chooses[i])
	}

	Engine.GET("/shopChoose", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		choose := DaoModle.ShopChoose{}
		err := context.ShouldBindQuery(&choose)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		value, ok := chooseMap.Load(choose.ShopId)
		if !ok {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     NOTFIND,
			})
			return
		}
		context.JSON(http.StatusOK, value.(*DaoModle.ShopChoose))
	})

	accessOrigin("/shopChoose")
}
