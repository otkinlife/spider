package doSite

import (
	"../config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type SiteHuaBan struct {
	Site
	url string
	c chan string
}

var imgDir string

func (s *SiteHuaBan) Download() {
	subImgDir = strconv.FormatInt(time.Now().Unix(), 10)
	urlList := s.getImgUrls()
	if urlList == nil || len(urlList) == 0 {
		panic("获取不到图片")
	}

	fmt.Println("总抓取图片数量：", len(urlList))
	//检查目录是否存在
	imgDir = outDir + subImgDir + "/"
	fmt.Println("图片保存路径：", imgDir)
	file, err := os.Stat(imgDir)
	if err != nil || !file.IsDir() {
		err := os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			panic("创建文件夹失败")
		}
	}

	i := 0
	for _, imgUrl := range urlList {
		i++
		go s.downloadImg(imgUrl, i)
		report := <- s.c
		fmt.Println(report)
	}
}

func (s *SiteHuaBan) getImgUrls() []string {
	var urlList []string
	res, err := http.Get(s.url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("目标地址请求失败")
	}
	htmlBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	htmlCode := string(htmlBytes)
	//使用正则匹配图片key
	reg := regexp.MustCompile(s.Site.Reg[1])
	keys := reg.FindAllStringSubmatch(htmlCode, -1)
	var url string
	for i := 0; i < len(keys); i++ {
		key := keys[i][1]
		// 过滤掉非图片类型的key
		if len(key) < 46 {
			continue
		}
		url = "https://hbimg.huabanimg.com/" + key
		urlList = append(urlList, url)
	}
	return urlList
}

func (s *SiteHuaBan) downloadImg(url string, preNo int) {
	config.WG.Add(1)
	prefix := strconv.Itoa(preNo) + "_"
	res := saveImages(url, imgDir, prefix)
	var str string
	if res {
		str = prefix + "下载成功"
	} else {
		str = prefix + "下载失败"
	}
	s.c <- str
	config.WG.Done()
}
