package doSite

import (
	"../config"
)

type SiteHuaBan struct {
	Site
	url string
}

func (s *SiteHuaBan) Download() {
	urlList := s.getImgUrls()
	if urlList == nil {
		panic("get img url failed")
	}
	for _, imgUrl := range urlList {
		config.WG.Add(1)
		go s.downloadImg(imgUrl)
		config.WG.Done()
	}
}

func (s *SiteHuaBan) getImgUrls() []string {
	//TODO:
	return []string{""}
}

func (s *SiteHuaBan) downloadImg(url string) {
	//TODO:
}
