package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func MyCron() {
	fmt.Println("启动定时任务:....")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("30 5 8 * * *", func() {
		fmt.Printf("运行定时任务-获取exchangerate-api数据:%v\n", time.Now())
		GetData() //获取最新数据-汇率 每天更新一次 8:05:30
	})
	c.AddFunc("5 */2 * * * *", func() {
		fmt.Printf("运行定时任务-获取币安数据:%v\n", time.Now())
		GetBiAnce() //获取币安数据
	})
	c.AddFunc("50 */8 * * * *", func() {
		fmt.Printf("运行定时任务-google sheets:%v\n", time.Now())
		GetRateFromGoogleSheets() //获取google sheets数据 每8分钟更新一次
	})
	c.Start()
	fmt.Println("定时任务启动成功!....")
}
