package routers

import (
	"github.com/AnCheTeng/govel/config"
	"github.com/AnCheTeng/govel/crawler"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getVendorCrawlers() []crawler.ICrawler {
	var vendorCrawlers []crawler.ICrawler

	log.Info(config.DefaultConfig)
	for _, vendor := range config.DefaultConfig.Vendors {
		var novels []crawler.Novel
		for _, n := range vendor.Novels {
			novels = append(novels, crawler.Novel{Name: n.Name, IndexURL: n.IndexURL})
		}
		if vendor.Name == "piaotian" {
			vendorCrawler := &crawler.PiaoTianCrawler{
				Hostname: vendor.Hostname,
			}
			vendorCrawler.AddNovels(novels)
			vendorCrawlers = append(vendorCrawlers, vendorCrawler)
		}
	}
	return vendorCrawlers
}

// GetLastUpdate - get last update by vendors
func GetLastUpdate(c *gin.Context) {
	vendorCrawlers := getVendorCrawlers()
	vendorCrawler := vendorCrawlers[0]
	ret := []gin.H{}
	for _, novel := range vendorCrawler.GetNovels() {
		chapterName, chapterURL := vendorCrawler.LatestChapter(novel)
		ret = append(ret, gin.H{
			"name":    novel.Name,
			"chapter": chapterName,
			"url":     chapterURL,
		})
		log.WithFields(log.Fields{
			"name": novel.Name, "chapterName": chapterName, "chapterURL": chapterURL,
		}).Info("Fetch success, latest chapter: ", chapterName)
	}
	c.JSON(200, ret)
}
