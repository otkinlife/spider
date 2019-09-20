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
	if url == "" {
		fmt.Println("地址不能为空")
		return
	}
	fmt.Println(config.ImgDir)
	fmt.Println("目标地址：", url)
	fmt.Println("任务开始：")
	doSite.DownloadImg(url)
	config.WG.Wait()
	fmt.Println("\n任务结束")
}
