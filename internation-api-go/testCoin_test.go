package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

// 测试从coinmarketcap中获取的货币信息，是否在币安中能得到价格
var jsonFile string = "./coinFile/401.json"

// 定义货币结构体
type Coin struct {
	Name   string `json:"name"`
	Symbol string `json:"code"`
	Logo   string `json:"logo"`
	//这一列是币安的状态，是否能从币安获取价格 默认值为true
	BianStatus bool `json:"bian"`
}

var badCoinsSymbol []string

func TestCoin(t *testing.T) {
	// 读取json文件,并解析为结构体
	var coins []Coin
	// 读取json文件
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	// 解析json文件
	err = json.Unmarshal(file, &coins)
	if err != nil {
		panic(err)
	}
	// 遍历结构体
	for i, coin := range coins {
		price, err := GetCoinPrice(coin)
		if err != nil {
			//记录下币种名称
			println("这种币种从币安上获取价格失败:", coin.Symbol)
			badCoinsSymbol = append(badCoinsSymbol, coin.Symbol)
			// 设置该币种的状态为false
			coins[i].BianStatus = false
			//睡500ms，防止请求过快
			time.Sleep(400 * time.Millisecond)
			continue
		} else {
			println("获取到价格:", coin.Symbol, ":", price)
			// 设置该币种的状态为true
			coins[i].BianStatus = true
		}
	}

	//把更新好的结构体，写入新的json文件:1a.json
	file, err = json.MarshalIndent(coins, "", " ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./coinFile/401a.json", file, 0644)
	if err != nil {
		panic(err)
	}
	// 测试结束

}

func GetCoinPrice(c Coin) (price string, err error) {
	queryString := c.Symbol + "USDT"
	url := "https://data-api.binance.vision/api/v3/ticker/price?symbol=" + queryString
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//解析到结构体
	var respType respType
	err = json.NewDecoder(resp.Body).Decode(&respType)
	if err != nil {
		return "", err
	}
	// 判断货币价格是否为空
	if respType.Price == "" {
		//返回错误信息
		return "", fmt.Errorf(c.Symbol + "价格为空")
	}
	return respType.Price, nil
}

type respType struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
