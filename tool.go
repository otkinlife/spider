package main

import (
	"./config"
	"./doSite"
	"flag"
	"fmt"
)

var url string
func init() {
	//初始化目标地址和输出目录
	flag.StringVar(&url, "url", "", "抓取目标地址")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("./spider --url=xxxx")
	}
}

func main() {
	fmt.Println("启动客户端")
	fmt.Println("Ctrl C Quit")
	for {
		url = ""
		fmt.Println("请输入要抓取的地址...")
		_, _ = fmt.Scanln(&url)
		if url == "" {
			fmt.Println("url不能为空")
			continue
		}
		fmt.Println("目标地址：", url)
		fmt.Println("任务开始：")
		doSite.DownloadImg(url)
		config.WG.Wait()
	}
	fmt.Println("\n任务结束")
}