package routers

import (
	"sync"

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

func goCrawlerWorker(
	novel crawler.Novel, vendorCrawler crawler.ICrawler,
	jobqueue chan gin.H, wg *sync.WaitGroup,
) {
	defer wg.Done()
	chapterName, chapterURL := vendorCrawler.LatestChapter(novel)
	jobqueue <- gin.H{
		"name":    novel.Name,
		"chapter": chapterName,
		"url":     chapterURL,
	}
	log.WithFields(log.Fields{
		"name": novel.Name, "chapterName": chapterName, "chapterURL": chapterURL,
	}).Info("Fetch success, latest chapter: ", chapterName)
}

// GetLastUpdate - get last update by vendors
func GetLastUpdate(c *gin.Context) {
	vendorCrawlers := getVendorCrawlers()
	vendorCrawler := vendorCrawlers[0]

	novels := vendorCrawler.GetNovels()
	totalNovels := len(novels)
	wg := sync.WaitGroup{}
	wg.Add(totalNovels)
	jobqueue := make(chan gin.H, totalNovels)
	for _, novel := range novels {
		go goCrawlerWorker(novel, vendorCrawler, jobqueue, &wg)
	}
	wg.Wait()

	ret := []gin.H{}
	for i := 0; i < totalNovels; i++ {
		ret = append(ret, <-jobqueue)
	}

	c.JSON(200, ret)
}
