package main

import (
	"./config"
	"./doSite"
	"flag"
	"fmt"
	"os"
)

var url string
var outDir string

func init() {
	//初始化目标地址和输出目录
	flag.StringVar(&url, "url", "", "抓取目标地址")
	flag.StringVar(&outDir, "out", "", "输出目录")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("./spider --url=xxxx")
	}
}

func main() {
	if url == "" {
		client()
	} else {
		if outDir == "" {
			outDir = config.ImgDir
		}
		file, err := os.Stat(outDir)
		if err != nil || !file.IsDir() {
			err := os.Mkdir(outDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		doSite.DownloadImg(url, outDir)
		config.WG.Wait()
	}
}

func client() {
	fmt.Println("启动客户端")
	fmt.Println("Ctrl C Quit")
	fmt.Println("请输入图片保存目录...(默认当前目录)")
	_, _ = fmt.Scanln(&outDir)
	if outDir == "" {
		outDir = config.ImgDir
	}
	file, err := os.Stat(outDir)
	if err != nil || !file.IsDir() {
		err := os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
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
		doSite.DownloadImg(url, outDir)
		config.WG.Wait()
		fmt.Println("\n任务结束")
	}
}
