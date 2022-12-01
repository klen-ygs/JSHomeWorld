package Serverlves

import (
	"GOServer/Config"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"sync"
)
import _ "github.com/go-sql-driver/mysql"
import "github.com/jinzhu/gorm"

var Engine *gin.Engine
var DB *gorm.DB

var CORS = Config.GetString("CORS")
var StaticURL = Config.GetString("StaticURL")

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
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	})

}

func getPhone(ctx *gin.Context, value interface{}) {
	defer func() {
		recover()
		fmt.Println("用户未登录")
	}()
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

var DataBaseUser = Config.GetString("DataBaseUser")
var Password = Config.GetString("DataBasePassword")
var DataBaseURL = Config.GetString("DataBaseURL")
var DataBase = Config.GetString("DataBase")

func init() {
	gin.SetMode("release")
	Engine = gin.Default()
	err := Engine.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", DataBaseUser+":"+Password+"@("+DataBaseURL+")/"+DataBase+"?charset=utf8mb4")
	DB = db
	if err != nil {
		panic(err)
	}
}
