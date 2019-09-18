package doSite

import (
	"../config"
	"fmt"
	"net/url"
)

//根据目标地址获取对应配置
func DownloadImg(targetUrl string) {
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
	default:
		fmt.Println("未匹配站点，使用默认方式")
		dObj := SiteDeFault{
			siteObj,
			targetUrl,
		}
		dObj.Download()
	}
}
