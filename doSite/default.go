package doSite

import (
	"../config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type SiteDeFault struct {
	Site
	url string
}

//解析并下载图片
func (s *SiteDeFault) Download() {
	urlList := s.getImgUrls()
	if urlList == nil || len(urlList) == 0 {
		panic("获取不到图片")
	}
	fmt.Println("总抓取图片数量：", len(urlList))

	//检查目录是否存在
	imgDir = config.ImgDir + subImgDir + "/"
	file, err := os.Stat(imgDir)
	if err != nil || !file.IsDir() {
		err := os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			panic("创建文件夹失败")
		}
	}

	//需要计算图片序号
	i := 0
	for _, imgUrl := range urlList {
		i++
		go s.downloadImg(imgUrl, i)
	}
}

//从页面解析<img>标签，并返回url列表
func (s *SiteDeFault) getImgUrls() []string {
	var urlList []string
	res, err := http.Get(s.url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("目标地址请求失败")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	subImgDir = doc.Find("title").Text()
	//处理url对象
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		url, _ := s.Attr("src")
		re := regexp.MustCompile(`^(https|http).*`)
		if re.MatchString(url) {
			urlList = append(urlList, url)
		}
	})
	return urlList
}

//下载图片
func (s *SiteDeFault) downloadImg(url string, preNo int) {
	config.WG.Add(1)
	prefix := strconv.Itoa(preNo) + "_"
	saveImages(url, imgDir, prefix)
	config.WG.Done()
}
