package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// 从google sheets中获取汇率
func GetRateFromGoogleSheets() {
	fmt.Println("开始获取google sheets数据")
	// proxyURL, err := url.Parse("http://localhost:1080")
	// if err != nil {
	// 	panic(err)
	// }

	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyURL),
	// }

	// client := &http.Client{
	// 	Transport: transport,
	// }

	// 发送 HTTP GET 请求
	resp, err := http.Get("https://docs.google.com/spreadsheets/d/e/2PACX-1vQlmuo1VKRfqHEmzRslTHf1xxAukaMo8TK3kNM7FUGqAhfjH16RoqUb0gJOf71oOzEfaJhQVx6SsyON/pub?gid=0&single=true&output=csv")
	if err != nil {
		fmt.Println("Error fetching the CSV file 1:", err)
		return
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP status code 2", resp.StatusCode)
		return
	}

	// 创建 CSV 读取器
	reader := csv.NewReader(resp.Body)

	// 读取所有行
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data 3:", err)
		return
	}
	// ExchangeRate 结构体表示汇率数据
	type ExchangeRate struct {
		Rate   float64
		Symbol string
	}
	// 解析 CSV 数据到结构体
	var exchangeRates []ExchangeRate
	for _, record := range records {
		// 解析汇率
		rate, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Println("Error parsing rate:", err)
			continue
		}

		// 创建 ExchangeRate 结构体
		exchangeRate := ExchangeRate{
			Rate:   rate,
			Symbol: record[1],
		}

		// 添加到切片
		exchangeRates = append(exchangeRates, exchangeRate)
	}
	Lock.Lock()         // 获取写锁
	defer Lock.Unlock() // 释放写锁，延迟到函数结束
	// 更新全局CurrencyInfoMap
	for _, er := range exchangeRates {
		if value, ok := CurrencyInfoMap[er.Symbol]; !ok {
			fmt.Println("From google Sheets:", er.Symbol, "不在CurrencyInfoMap中")
		} else {
			value.ExchangeRate = er.Rate
			value.UpdateTime = "g" + time.Now().Local().Format("2006-01-02 15:04:05")
			CurrencyInfoMap[er.Symbol] = value
			// fmt.Println("From google Sheets:", er.Symbol, "更新成功:", "旧汇率:", r, "新汇率:", er.Rate)
		}
	}
	CurrencyInfoList = make([]CurrencyInfo, 0, cap(CurrencyInfoList))
	for _, value := range CurrencyInfoMap {
		if value.ExchangeRate > 0 {
			CurrencyInfoList = append(CurrencyInfoList, value)

		}
	}
	fmt.Println("google sheets数据更新完毕")
}
