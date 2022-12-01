package Servers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func returnIndexHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func init() {
	Engine.LoadHTMLFiles("..\\vueshop\\dist\\index.html")
	Engine.Static("/js", "..\\vueshop\\dist\\js")
	Engine.Static("/css", "..\\vueshop\\dist\\css")
	Engine.Static("/img", "..\\vueshop\\dist\\img")
	Engine.Static("/static", "..\\vueshop\\dist\\static")
	Engine.StaticFile("/favicon.ico", "..\\vueshop\\dist\\static\\favicon.ico")
	Engine.GET("/", returnIndexHtml)
	Engine.GET("/select", returnIndexHtml)
	Engine.GET("/Shop", returnIndexHtml)
	Engine.GET("/AddShop", returnIndexHtml)
	Engine.GET("/ShopList", returnIndexHtml)
	Engine.GET("/selected", returnIndexHtml)
}
