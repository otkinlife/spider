package doSite

import (
	"../config"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Site struct {
	Site int
	Host string
	Reg  []string
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

//获取对应站点对象
func GetTypeObj(host string) Site {
	siteType := config.UrlConvert[host]
	siteObj := Site{
		siteType,
		host,
		config.SiteReg[siteType],
	}
	return siteObj

}

//下载图片
func saveImages(imgUrl string, dir string, prefix string) {
	//去掉最左边的'/'
	filename := dir + prefix + RandString(16) + ".jpg"

	exists := checkExists(filename)
	if exists {
		return
	}
	//获取静态页面
	response, err := http.Get(imgUrl)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", imgUrl, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}
	defer image.Close()
	_, _ = image.Write(data)
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
