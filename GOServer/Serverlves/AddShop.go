package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func init() {
	var mutex = sync.Mutex{}

	Engine.POST("/addShop", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		newShop := DaoModle.Shop{}
		err := context.BindJSON(&newShop)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
			return
		}
		func() {
			mutex.Lock()
			defer mutex.Unlock()
			DB.Create(&newShop)
			DB.Select("id").Last(&newShop)
		}()
		context.JSON(http.StatusOK, gin.H{
			"id":  newShop.Id,
			"Err": NONE,
		})
	})

	accessOrigin("/addShop")
}
