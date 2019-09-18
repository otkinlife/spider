package config

const DEFAULT  = 0
const HUABAN  = 1

var ImgDir = "./img/"

//域名和站点配置
var UrlConvert = map[string]int{
	"huaban.com": HUABAN,
}

var SiteReg = map[int][]string {
	HUABAN: {
		`^(https?://hbimg\.[a-zA-Z0-9]{2,5}\.upaiyun\.com/.*?)_[a-zA-Z0-9]{2,8}$`,
		`^(https?://img\.hb\.aicdn\.com/.*?)_[a-zA-Z0-9]{2,8}$`,
		`^(https?://hbimg\.huabanimg\.com/.*?)_[a-zA-Z0-9]{2,8}$`,
	},
	DEFAULT:{

	},
}
