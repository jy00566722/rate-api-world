package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetConfig() {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	viper.WatchConfig()
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			fmt.Println("no such config file")
		} else {

			fmt.Println("read config error")
		}
		fmt.Println(err)
	}

	err = viper.UnmarshalKey("crypto", &CryptoList)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	for _, crypto := range CryptoList {
		CryptoMap[crypto.Query] = crypto
		// fmt.Printf("Code: %s, Query: %s, Name: %s, FlagURL: %s\n",
		// 	crypto.Code, crypto.Query, crypto.Name, crypto.FlagURL)
	}
	// fmt.Printf("CryptoMap: %v\n", CryptoMap)

}
