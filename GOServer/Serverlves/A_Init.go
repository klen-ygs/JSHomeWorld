package Serverlves

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"sync"
)
import _ "github.com/go-sql-driver/mysql"
import "github.com/jinzhu/gorm"

var Engine *gin.Engine
var DB *gorm.DB

const AJAXOrigin = "http://localhost:8080"

const (
	PHONEERR     = "PhoneErr"
	PASSWORDERR  = "PasswordErr"
	NONE         = ""
	MODLEERR     = "ModelErr"
	NOTFIND      = "NotFind"
	INSERTERR    = "InsertErr"
	VIDEOPATHERR = "VideoPathErr"
)

var (
	phoneMap   sync.Map
	loginCount uint64 = 1
	countMap   sync.Map
)

// response 处理回复消息
type response struct {
	Request bool
	Err     string
}

// 函数功能：设置axios嗅探请求
func accessOrigin(url string) {
	Engine.OPTIONS(url, func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,XFILENAME,XFILECATEGORY,XFILESIZE")
		context.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "true")
	})

}

func getPhone(ctx *gin.Context, value interface{}) {
	elem := reflect.ValueOf(value).Elem()
	phone := elem.FieldByName("Phone")
	if phone.String() != "" {
		return
	}
	cookie, err := ctx.Cookie("login")
	if err != nil {
		return
	}
	parseUint, _ := strconv.ParseUint(cookie, 10, 64)
	load, _ := phoneMap.Load(parseUint)
	phone.SetString(load.(string))
}

func init() {
	Engine = gin.Default()
	db, err := gorm.Open("mysql", "root:28885/YGSnizhan@(127.0.0.1:3306)/vue_shop?charset=utf8mb4")
	DB = db
	if err != nil {
		panic(err)
	}
}
