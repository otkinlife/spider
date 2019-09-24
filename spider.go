package main

import (
	"./config"
	"./doSite"
	"flag"
	"fmt"
	"os"
)

var mode string
var url string
var outDir string

func init() {
	//初始化目标地址和输出目录
	flag.StringVar(&mode, "m", "client", "启动模式,默认client")
	flag.StringVar(&url, "u", "", "抓取目标地址")
	flag.StringVar(&outDir, "o", "", "输出目录")
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("./spider -m=script -u=http://xxxx.com -p=/xxx/xxx/")
	}
}

func main() {
	fmt.Println(mode)
	switch mode {
	case "daemon":
		daemon()
		break
	case "script":
		script()
		break
	case "client":
		client()
		break
	}
}

//启动客户端模式
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

//脚本模式
func script() {
	if outDir == "" {
		outDir = config.ImgDir
	}
	fmt.Println(outDir)
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

//守护进程模式
func daemon() {
	fmt.Println("努力实现....")
	//TODO：实现守护进程，类似于一个队列
}
