package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url, filepath string) error {
	// 发起 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建文件来保存下载的内容
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 将下载的内容保存到文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("文件 %s 下载完成\n", filepath)
	return nil
}

func gettime() {
	// 获取当前时间
	currentTime := time.Now()

	// 格式化时间为字符串
	timeString := currentTime.Format("2006-01-02 15:04:05") + "\n"

	// 打开文件以追加文本
	filePath := "update_time.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件时出错:", err)
		return
	}
	defer file.Close()

	// 将时间字符串写入文件
	_, err = file.WriteString(timeString)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	fmt.Printf("当前时间已追加到文件 %s\n", filePath)
}

func main() {
	gettime()

	fileURL := "https://iitzh.com/cn.m3u"
	filePath := "cn.m3u"

	err := downloadFile(fileURL, filePath)
	if err != nil {
		fmt.Println("下载文件时出错:", err)
	}
	
}
