package doSite

import (
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
	//根据域名获取对应站点的对象
	dObj := GetTypeObj(u.Host, targetUrl)
	dObj.Download()
}


