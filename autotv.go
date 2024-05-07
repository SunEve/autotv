package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

func addxmlpath(addString, filepath string) {
	// 读取原始文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件时出错:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	// 在第一行末尾添加指定字符串
	if len(lines) > 0 {
		lines[0] += addString
	}

	// 将修改后的内容写回文件
	output := []byte(strings.Join(lines, "\n"))
	err = ioutil.WriteFile(filePath, output, 0644)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	fmt.Printf("指定字符串已成功添加到文件 %s 的第一行末尾\n", filePath)
}

func main() {
	gettime()

	fileURL := "https://cdn.jsdelivr.net/gh/BurningC4/Chinese-IPTV@master/TV-IPV4.m3u"
	filePath := "public/cn.m3u"

	err := downloadFile(fileURL, filePath)
	if err != nil {
		fmt.Println("下载文件时出错:", err)
		return
	}
	
	filePath1 := "public/cn11.m3u"
	addString := " x-tvg-url=\"https://live.fanmingming.com/e.xml\""
	addxmlpath(addString, filepath1)
}
