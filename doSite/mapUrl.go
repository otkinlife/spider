package doSite

import (
	"../config"
	"fmt"
	"net/url"
)

var outDir = ""
var subImgDir = ""

//根据目标地址获取对应配置
func DownloadImg(targetUrl string, dir string) {
	outDir = dir
	//解析并获取输入地址的域名
	u, err := url.Parse(targetUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Host)
	siteObj := GetTypeObj(u.Host)
	switch siteObj.Site {
	case config.HUABAN:
		fmt.Println("检测为花瓣网")
		dObj := SiteHuaBan{
			siteObj,
			targetUrl,
		}
		dObj.Download()
	case config.TUCHONG:
		fmt.Printf("检测为图虫网")
		dObj := SiteTuChong{
			siteObj,
			targetUrl,
		}
		dObj.Download()
	default:
		fmt.Println("未匹配站点，使用默认方式")
		dObj := SiteDeFault{
			siteObj,
			targetUrl,
		}
		dObj.Download()
	}
}
