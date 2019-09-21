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

type SiteTuChong struct {
	Site
	url string
}

func (s *SiteTuChong) Download() {
	subImgDir = strconv.FormatInt(time.Now().Unix(), 10)
	urlList := s.getImgUrls()
	if urlList == nil || len(urlList) == 0 {
		panic("获取不到图片")
	}

	fmt.Println("总抓取图片数量：", len(urlList))
	//检查目录是否存在
	imgDir = config.ImgDir + subImgDir + "/"
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
	}
}

func (s *SiteTuChong) getImgUrls() []string {
	var urlList []string
	client := &http.Client{}
	//生成要访问的url
	//提交请求
	request, err := http.NewRequest("GET", s.url, nil)

	//增加header选项
	request.Header.Add("Cookie", "lang=zh;")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	res, _ := client.Do(request)
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
	reg := regexp.MustCompile(s.Site.Reg[0])
	keys := reg.FindAllStringSubmatch(htmlCode, -1)
	var url string
	for i := 0; i < len(keys); i++ {
		key := keys[i][1]
		url = "http://icweiliimg6.pstatp.com/weili/l/" + key + ".webp"
		urlList = append(urlList, url)
	}
	return urlList
}

func (s *SiteTuChong) downloadImg(url string, preNo int) {
	config.WG.Add(1)
	prefix := strconv.Itoa(preNo) + "_"
	saveImages(url, imgDir, prefix)
	config.WG.Done()
}
