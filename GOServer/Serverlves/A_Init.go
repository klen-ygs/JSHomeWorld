package Serverlves

import "github.com/gin-gonic/gin"
import _ "github.com/go-sql-driver/mysql"
import "github.com/jinzhu/gorm"

var Engine *gin.Engine
var DB *gorm.DB

const (
	PHONEERR    = "PhoneErr"
	PASSWORDERR = "PasswordErr"
	NONE        = ""
	MODLEERR    = "ModelErr"
	NOTFIND     = "NotFind"
)

// response 处理回复消息
type response struct {
	request bool
	Err     string
}

// 函数功能：设置axios嗅探请求
func accessOrigin(url string) {
	Engine.OPTIONS(url, func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,XFILENAME,XFILECATEGORY,XFILESIZE")
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	})

}

func init() {
	Engine = gin.Default()
	db, err := gorm.Open("mysql", "root:28885/YGSnizhan@(127.0.0.1:3306)/vue_shop")
	DB = db
	if err != nil {
		panic(err)
	}
}
