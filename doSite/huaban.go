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
}

func (s *SiteHuaBan) Download() {
	subImgDir = strconv.FormatInt(time.Now().Unix(), 10)
	urlList := s.getImgUrls()
	if urlList == nil {
		panic("get img url failed")
	}
	fmt.Println("总抓取图片数量：", len(urlList))
	//检查目录是否存在
	imgDir = imgDir + subImgDir + "/"
	fmt.Println(imgDir)
	file, err := os.Stat(imgDir)
	if err != nil || !file.IsDir() {
		err := os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			panic("create dir failed")
		}
	}

	i := 0
	for _, imgUrl := range urlList {
		i++
		go s.downloadImg(imgUrl, i)
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
		panic("status code error")
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
	saveImages(url, imgDir, prefix)
	config.WG.Done()
}