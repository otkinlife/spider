package doSite

type SiteDownload interface {
	//下载图片入口方法
	Download(url string)
	//获取要下载的图片列表方法
	getImgUrls(url string) []string
	//真正的下载图片方法
	downloadImg(url string)
}
