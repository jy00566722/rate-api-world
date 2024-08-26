package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var CryptoList []CoinInfo
var CryptoMap = make(map[string]CoinInfo)

func GetBiAnce() {
	data := GetBiAnceData()
	VirtualCurrencyInfoList_temp := []CurrencyInfo{}
	for _, v := range data {
		cryptoinfo := CryptoMap[v.Symbol]
		//价格转换float64
		price, _ := strconv.ParseFloat(v.Price, 64)

		CurrencyInfo_ := CurrencyInfo{
			Code:         cryptoinfo.Code,
			ExchangeRate: price,
			Country:      cryptoinfo.Code,
			Currency:     cryptoinfo.Name,
			Country_code: cryptoinfo.Name,
			IsVirtual:    true,
			FlagURL:      cryptoinfo.FlagURL,
		}
		VirtualCurrencyInfoList_temp = append(VirtualCurrencyInfoList_temp, CurrencyInfo_)

	}
	usdt := CurrencyInfo{
		Code:         "USDT",
		ExchangeRate: 1,
		Country:      "USDT",
		Currency:     "Tether USDt",
		Country_code: "us",
		IsVirtual:    true,
		FlagURL:      "https://svg.oeo.li/svg/USDT.svg",
	}
	VirtualCurrencyInfoList_temp = append(VirtualCurrencyInfoList_temp, usdt)
	Lock.Lock()
	defer Lock.Unlock()
	VirtualCurrencyInfoList = VirtualCurrencyInfoList_temp
	TimeLastGet_crypto = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("币安数据获取完")
}

// GetData函数，用来获取币安接口返回的数据
func GetBiAnceData() []respDateType {

	// 创建一个字符串切片来存储所有的Query
	var queries []string
	for _, crypto := range CryptoList {
		queries = append(queries, fmt.Sprintf("%q", crypto.Query))
	}

	// 将字符串切片连接成一个字符串
	queryString := fmt.Sprintf("[%s]", strings.Join(queries, ","))

	url := "https://data-api.binance.vision/api/v3/ticker/price?symbols=" + queryString
	// fmt.Printf("url: %v\n", url)
	//定义一个map类型的变量，用来存储币的名称和价格
	var coinMap []respDateType = GetBinancePrice(url)

	return coinMap
}

// 实现GetBinancePrice函数
func GetBinancePrice(url string) []respDateType {
	//发送HTTP请求，获取币安接口返回的数据
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//解析JSON数据
	var data []respDateType
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	// //创建一个map，用来存储币的名称和价格
	// coinMap := make(map[string]string)
	// for _, d := range data {
	// 	coinMap[d.Symbol] = d.Price
	// }
	//返回币的名称和价格
	return data
}

// 定义加密货币信息
type CoinInfo struct {
	Code    string `mapstructure:"code" json:"code"`
	Query   string `mapstructure:"query" json:"query"`
	Name    string `mapstructure:"name" json:"name"`
	FlagURL string `mapstructure:"flagUrl" json:"flagURL"`
}

type respDateType struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
