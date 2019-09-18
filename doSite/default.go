package doSite

import (
	"../config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"regexp"
)

var imgDir = config.ImgDir

type SiteDeFault struct {
	Site
	url string
}

//解析并下载图片
func (s *SiteDeFault)Download() {
	urlList := s.getImgUrls()
	if urlList == nil {
		panic("get img url failed")
	}
	fmt.Println("总抓取图片数量：", len(urlList))
	i := 1;
	for _, imgUrl := range urlList {
		fmt.Printf("\r当前进度：%d", i)
		i++
		go s.downloadImg(imgUrl)
	}
}

//从页面解析<img>标签，并返回url列表
func (s *SiteDeFault)getImgUrls() []string {
	var urlList []string
	res, err := http.Get(s.url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	// Find the review items
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		url, _ := s.Attr("src")
		re := regexp.MustCompile(`^(https|http).*`)
		if (re.MatchString(url)) {
			urlList = append(urlList, url)
		}
	})
	return urlList
}

//下载图片
func (s *SiteDeFault)downloadImg(url string) {
	config.WG.Add(1)
	//检查目录是否存在
	file, err := os.Stat(imgDir)
	if err != nil || !file.IsDir() {
		err := os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			panic("create dir failed")
		}
	}
	saveImages(url, imgDir)
	config.WG.Done()
}

