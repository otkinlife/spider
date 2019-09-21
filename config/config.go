package config

//默认
const DEFAULT = 0

//花瓣网
const HUABAN = 1

//图虫网
const TUCHONG = 2

var ImgDir = "./img/"

//域名和站点配置
var UrlConvert = map[string]int{
	"huaban.com":        HUABAN,
	"stock.tuchong.com": TUCHONG,
	"tuchong.com":       TUCHONG,
}

//每个站点所用的正则
var SiteReg = map[int][]string{
	HUABAN: {
		`"pin_id":(\d+),`,
		`"key":"(.*?)"`,
	},
	TUCHONG: {
		`"imageId":"(.*?)",`,
	},
	DEFAULT: {},
}
