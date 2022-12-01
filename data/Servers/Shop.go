package Servers

import (
	"container/heap"
	"data/DaoModle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"sync"
)

type getShopRequest struct {
	ShopId     uint64
	MinId      uint64
	MaxId      uint64
	FindMethod string
	Page       int
	SearchWord string
}

var ShopMap sync.Map
var shops DaoModle.ShopSlice

func getFirstPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, shops[:15])
}

func getFirstBySearch(searchWord string) DaoModle.ShopSlice {
	tarArr := make(DaoModle.ShopSlice, 0)
	for i, _ := range shops {
		if strings.Index(shops[i].ShopTitleText, searchWord) != -1 {
			tarArr = append(tarArr, shops[i])
		} else if searchWord == "" {
			tarArr = append(tarArr, shops[i])
		}
		if len(tarArr) == 15 {
			break
		}
	}
	return tarArr
}

func findNext(id uint64, searchText string) DaoModle.ShopSlice {
	var idx = 0
	for i, _ := range shops { //找到首位id
		if shops[i].Id == id {
			idx = i
			break
		}
	}
	lastIdx := idx
	tarArr := make(DaoModle.ShopSlice, 0)
	for ; lastIdx < len(shops) && len(tarArr) < 15; lastIdx++ {
		if strings.IndexAny(shops[lastIdx].ShopTitleText, searchText) != -1 {
			tarArr = append(tarArr, shops[lastIdx])
		} else if searchText == "" {
			tarArr = append(tarArr, shops[lastIdx])
		}
	}
	return tarArr
}

func findLast(id uint64, searchText string) DaoModle.ShopSlice {
	tarArr := make(DaoModle.ShopSlice, 0)
	for i := 0; id != shops[i].Id; i++ {
		if strings.IndexAny(shops[i].ShopTitleText, searchText) != -1 {
			tarArr = append(tarArr, shops[i])
		} else if searchText == "" {
			tarArr = append(tarArr, shops[i])
		}
		if len(tarArr) > 15 {
			tarArr = tarArr[1:]
		}
	}
	return tarArr
}

func init() {
	file, err := os.ReadFile("./shop.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &shops)
	if err != nil {
		panic(err)
	}
	heap.Init(&shops)
	for i, _ := range shops {
		ShopMap.Store(shops[i].Id, &shops[i])
	}

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
			value, ok := ShopMap.Load(request.ShopId)
			if !ok {
				context.JSON(http.StatusOK, response{
					Request: false,
					Err:     NOTFIND,
				})
				return
			}
			context.JSON(http.StatusOK, value.(*DaoModle.Shop))
			return
		}
		if request.SearchWord != "" {
			SearchChan <- request.SearchWord // 送给热点逻辑处理
		}
		var ShopArr []DaoModle.Shop
		if request.FindMethod == "Next" {
			ShopArr = findNext(request.MaxId, request.SearchWord)
		} else if request.FindMethod == "Last" {
			ShopArr = findLast(request.MinId, request.SearchWord)
		} else {
			ShopArr = getFirstBySearch(request.SearchWord)
		}
		context.JSON(http.StatusOK, &ShopArr)
	})

	accessOrigin("/getShop")

}
