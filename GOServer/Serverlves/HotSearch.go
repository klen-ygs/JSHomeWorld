package Serverlves

import (
	"GOServer/DaoModle"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var searchText chan string = make(chan string, 10)

func addSearchText() {
	var hot = DaoModle.HotSearch{}
	for {
		hot.SearchText = <-searchText
		hot.SearchCount = 0
		DB.Where("search_text = ?", hot.SearchText).FirstOrCreate(&hot)
		hot.SearchCount++
		DB.Model(&hot).Where("search_text = ?", hot.SearchText).Update(&hot)
	}
}

func init() {
	go addSearchText()

	//每十秒更新一次热搜榜
	var hotArr = make([]DaoModle.HotSearch, 10)
	HotUpdateTime := time.NewTicker(10 * time.Second)
	go func() {
		for {
			<-HotUpdateTime.C
			DB.Order(" search_count desc").Find(&hotArr)
		}
	}()

	Engine.GET("/hotSearch", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		context.JSON(http.StatusOK, &hotArr)
	})
}
