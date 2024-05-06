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

func main() {
	fileURL := "https://iitzh.com/cn.m3u"
	filePath := "cn.m3u"

	err := downloadFile(fileURL, filePath)
	if err != nil {
		fmt.Println("下载文件时出错:", err)
	}
}
