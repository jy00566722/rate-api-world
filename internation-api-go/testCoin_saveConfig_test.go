package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"testing"
	"time"
)

var GoodCoins []CoinInfo

func TestSaveCoinConfigFile(t *testing.T) {
	//从coinFile文件夹下的5个文件中读取数据:1a.json,101a.json,201a.json,301a.json,401a.json
	//获取bian为true的币种,并保存到coinFile文件夹下的coinConfig.json文件中,
	//如果coinConfig.json文件存在,则覆盖

	fileList := [2]string{"1a.json", "201a.json"}
	for _, v := range fileList {
		//读取json文件,获取币种信息
		file, err := os.ReadFile("./coinFile/" + v)
		if err != nil {
			t.Error(err)
		}
		var coins []Coin
		err = json.Unmarshal(file, &coins)
		if err != nil {
			t.Error(err)
		}
		//遍历币种信息,获取bian为true的币种，并保存到GoodCoins中
		for _, coin := range coins {
			if coin.BianStatus == true {
				GoodCoins = append(GoodCoins, CoinInfo{
					Code:    coin.Symbol,
					Query:   coin.Symbol + "USDT",
					Name:    coin.Name,
					FlagURL: coin.Logo,
				})
			}
		}

	}
	//将GoodCoins中的币种信息保存到coinConfig.json文件中
	file, err := json.MarshalIndent(GoodCoins, "", " ")
	if err != nil {
		t.Error(err)
	}
	//将数据写入coinConfig.json文件中
	err = os.WriteFile("./coinConfig-200.json", file, 0755)
	if err != nil {
		t.Error(err)
	}
	t.Log("coinConfig-200.json文件保存成功")

}

// 从coinFile文件夹下的5个文件中读取数据:1a.json,101a.json,201a.json,301a.json,401a.json
// 获取每个币种的logo,然后下载对应的log图片,并保存到coinFile文件夹下的coinLogo文件夹中
func TestSaveCoinLogo(t *testing.T) {
	//从coinFile文件夹下的5个文件中读取数据:1a.json,101a.json,201a.json,301a.json,401a.json
	fileList := [5]string{"1a.json", "101a.json", "201a.json", "301a.json", "401a.json"}
	for _, v := range fileList {
		//读取json文件,获取币种信息
		file, err := os.ReadFile("./coinFile/" + v)
		if err != nil {
			t.Error(err)
		}
		var coins []Coin
		err = json.Unmarshal(file, &coins)
		if err != nil {
			t.Error(err)
		} else {
			t.Log("读取币种信息成功")
		}
		//遍历币种信息,获取币种的logo,并下载
		for _, coin := range coins {
			filename, err := getFilenameFromURL(coin.Logo)
			if err != nil {
				fmt.Println("Error getting filename:", err, coin.Logo)
				continue
			}
			// 设置保存的目录和文件名
			savePath := path.Join("downloaded_images", filename)
			if err := os.MkdirAll("downloaded_images", 0755); err != nil {
				fmt.Println("Error creating directory:", err, savePath)
				continue
			}
			// 下载并保存文件
			if err := downloadFile(coin.Logo, savePath); err != nil {
				fmt.Println("Error downloading file:", err, coin.Logo)
				continue
			}

			fmt.Println("File downloaded successfully:", savePath)
			time.Sleep(time.Microsecond * 200)
		}
	}
}

// getFilenameFromURL 解析URL并返回文件名
func getFilenameFromURL(link string) (string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return path.Base(parsedURL.Path), nil
}

// downloadFile 下载文件并保存到指定路径
func downloadFile(URL, filepath string) error {
	// 发送GET请求
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %d %s", resp.StatusCode, resp.Status)
	}

	// 创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 将响应体复制到文件
	_, err = io.Copy(out, resp.Body)
	return err
}
