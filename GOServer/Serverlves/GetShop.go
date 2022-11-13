package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getShopRequest struct {
	ShopId     uint64
	MinId      uint64
	MaxId      uint64
	FindMethod string
	Page       int
	SearchWord string
}

func getFirstPage(context *gin.Context) {
	ShopArr := make([]DaoModle.Shop, 20)
	DB.Limit(20).Find(&ShopArr)
	context.JSON(200, &ShopArr)
}

func init() {
	Engine.GET("/getShop", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		request := getShopRequest{}
		err := context.BindQuery(&request)
		if err != nil {
			context.JSON(http.StatusOK, response{
				Request: false,
				Err:     MODLEERR,
			})
		}
		if request.SearchWord == "" && request.Page == 1 {
			getFirstPage(context)
			return
		} else if request.ShopId != 0 {
			shop := DaoModle.Shop{}
			DB.Model(&shop).Where("id = ?", request.ShopId).First(&shop)
			context.JSON(http.StatusOK, &shop)
			return
		}
		ShopArr := make([]DaoModle.Shop, 15)
		request.SearchWord = "%" + request.SearchWord + "%"
		if request.FindMethod == "Next" {
			DB.Where("shop_title_text like ? and id > ?", request.SearchWord, request.MaxId).Find(&ShopArr)
		} else if request.FindMethod == "Last" {
			DB.Where("shop_title_text like ? and id < ?", request.SearchWord, request.MinId).Find(&ShopArr)
		} else {
			DB.Where("shop_title_text like ?", request.SearchWord).Find(&ShopArr)
		}
		context.JSON(http.StatusOK, &ShopArr)
	})

	accessOrigin("/getShop")
}
