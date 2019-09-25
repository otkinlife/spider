package doSite

import (
	"../config"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Site struct {
	Site int
	Host string
	Reg  []string
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

//获取对应站点对象，并返回相关配置对象
func GetTypeObj(host string, targetUrl string) SiteDownload {
	siteType := config.UrlConvert[host]
	siteObj := Site{
		siteType,
		host,
		config.SiteReg[siteType],
	}
	c := make(chan string, 10)
	var dObj SiteDownload
	switch siteObj.Site {
	case config.HUABAN:
		fmt.Println("检测为花瓣网")
		dObj = &SiteHuaBan{
			siteObj,
			targetUrl,
			c,
		}
	case config.TUCHONG:
		fmt.Printf("检测为图虫网")
		dObj = &SiteTuChong{
			siteObj,
			targetUrl,
			c,
		}
	default:
		fmt.Println("未匹配站点，使用默认方式")
		dObj = &SiteDeFault{
			siteObj,
			targetUrl,
			c,
		}
	}
	return dObj
}

//下载图片
func saveImages(imgUrl string, dir string, prefix string) bool {
	//去掉最左边的'/'
	filename := dir + prefix + RandString(16) + ".jpg"

	exists := checkExists(filename)
	if exists {
		return true
	}
	//获取静态页面
	response, err := http.Get(imgUrl)
	if err != nil {
		log.Println("get img_url failed:", err)
		return false
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", imgUrl, err)
		return false
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return false
	}
	defer image.Close()
	_, _ = image.Write(data)
	return true
}

func GetSchedule(i int, sum int) string {
	var schedule float64
	schedule = float64(i) / float64(sum)
	schedule = math.Trunc(schedule * 100)
	fmt.Print("\r","下载进度：", schedule, "%")
	return "\r下载进度：" + strconv.Itoa(int(schedule)) + "%"
}

//检测文件是否存在
func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

//取随机名
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
