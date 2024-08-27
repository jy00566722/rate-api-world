package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
)

var CurrencyInfoList []CurrencyInfo
var VirtualCurrencyInfoList []CurrencyInfo
var CurrencyInfoMap = make(map[string]CurrencyInfo) //增加map，用更存放汇率信息，然后定时去转换为切片CurrencyInfoList ,这样方便更新汇率
var Lock sync.RWMutex
var qClient *qmgo.Client //mongodb的客户端
var err error
var TimeLastUpdate string     //汇率接口返回的最后数据时间
var TimeLastGet string        //程序去获取接口的最后时间
var TimeLastGet_crypto string //程序去获取加密货币价格的最后时间

func init() {
	GetConfig() //获取配置文件
	qClient, err = qmgo.NewClient(context.Background(), &qmgo.Config{Uri: viper.GetString("mongodbDns")})
	if err != nil {
		fmt.Printf("连接mongodb出错: %v\n", err)
	} else {
		fmt.Println("连接mongodb成功")
	}
}

func main() {
	// GetConfig() //获取配置文件
	InitMapData()             //初始化map,从本地配置文件中解析数据
	GetData()                 //获取最新数据-汇率,这是从汇率网站获取的,它的数据一天只更新一次
	GetBiAnce()               //获取币安数据
	GetRateFromGoogleSheets() //从谷歌表格中获取数据
	go MyCron()               //启动定时任务
	//准备gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	LoadFeedbackRouter(r)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word Currency!")
	})

	r.GET("/v1/exchange", func(c *gin.Context) {
		Lock.RLock()
		// dataCopy := make([]CurrencyInfo, len(CurrencyInfoList))
		// copy(dataCopy, CurrencyInfoList)
		defer Lock.RUnlock()
		c.Header("Data-Server", viper.GetString("DataServer"))
		c.JSON(http.StatusOK, gin.H{
			"code":           200,
			"base":           "USD",
			"data":           CurrencyInfoList,
			"TimeLastUpdate": TimeLastUpdate,
			"TimeLastGet":    TimeLastGet,
		})
	})
	r.GET("/v1/exchange_crypto", func(c *gin.Context) {
		Lock.RLock()
		defer Lock.RUnlock()
		c.Header("Data-Server", viper.GetString("DataServer"))
		c.JSON(http.StatusOK, gin.H{
			"code":                   200,
			"base":                   "USDT",
			"data":                   VirtualCurrencyInfoList,
			"TimeLastUpdate":         TimeLastGet_crypto,
			"TimeLastGet":            TimeLastGet_crypto,
			"periodInMinutes_crypto": viper.GetInt("periodInMinutes_crypto"),
		})
	})
	r.GET("/v1/message", func(c *gin.Context) {
		message := viper.Get("message")
		c.JSON(200, gin.H{"code": 20000, "message": "ok", "data": message})
	})
	r.Run(":8012")
}

//GOOS=linux GOARCH=amd64 go build -o internation-rate-1125a
