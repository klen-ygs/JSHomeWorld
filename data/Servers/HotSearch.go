package Servers

import (
	"container/heap"
	"data/DaoModle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
)

var SearchChan = make(chan string, 20)
var hotSearch DaoModle.HotSearchSlice
var searchMap = sync.Map{}

func init() {
	file, err := os.ReadFile("./hot_search.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &hotSearch)
	if err != nil {
		panic(err)
	}
	for i, _ := range hotSearch {
		searchMap.Store(hotSearch[i].SearchText, hotSearch[i].SearchCount)
	}
	heap.Init(&hotSearch)
	go func() {
		for {
			searchText := <-SearchChan
			value, ok := searchMap.Load(searchText)
			if ok {
				searchMap.Store(searchText, value.(uint64)+1)
			} else {
				heap.Push(&hotSearch, DaoModle.HotSearch{
					SearchCount: 0,
					SearchText:  searchText,
				})
				searchMap.Store(searchText, &hotSearch[len(hotSearch)-1])
			}
		}
	}()

	Engine.GET("/hotSearch", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		if len(hotSearch) < 10 {
			context.JSON(http.StatusOK, &hotSearch)
			return
		}
		context.JSON(http.StatusOK, hotSearch[:10])
	})
}
