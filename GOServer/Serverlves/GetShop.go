package Serverlves

import (
	"GOServer/DaoModle"
	"fmt"
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
	var ShopArr []DaoModle.Shop
	DB.Limit(15).Find(&ShopArr)
	context.JSON(200, &ShopArr)
}

func init() {
	Engine.GET("/getShop", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
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
		if request.SearchWord != "" {
			searchText <- request.SearchWord // 送给热点逻辑处理
		}
		var ShopArr []DaoModle.Shop
		request.SearchWord = "%" + request.SearchWord + "%"
		fmt.Println(request.FindMethod)
		if request.FindMethod == "Next" {
			DB.Where("shop_title_text like ? and id > ?", request.SearchWord, request.MaxId).Limit(15).Find(&ShopArr)
		} else if request.FindMethod == "Last" {
			DB.Where("shop_title_text like ?", request.SearchWord).Limit(15).Offset(15).Find(&ShopArr)
		} else {
			DB.Where("shop_title_text like ?", request.SearchWord).Limit(15).Find(&ShopArr)
		}
		context.JSON(http.StatusOK, &ShopArr)
	})

	accessOrigin("/getShop")
}
