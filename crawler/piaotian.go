package crawler

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

// PiaoTianCrawler - Get latest update from https://www.ptwxz.com/
type PiaoTianCrawler struct {
	Hostname string
	novels   []Novel
}

// AddNovels add novels
func (c *PiaoTianCrawler) AddNovels(novels []Novel) {
	c.novels = append(c.novels, novels...)
}

// GetNovels get novels
func (c *PiaoTianCrawler) GetNovels() []Novel {
	return c.novels
}

// LatestChapter - Get the latest chapter
func (c *PiaoTianCrawler) LatestChapter(novel Novel) (chapterName, chapterURL string) {
	// Request the HTML page.
	client := getClient()
	indexURL := c.Hostname + novel.IndexURL
	res, err := client.Get(indexURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".hottext").Each(
		func(i int, s *goquery.Selection) {
			if gbkToUTF8(s.Text()) == "最新章节：" {
				chapterName = gbkToUTF8(s.Next().Text())
				chapterURL = c.Hostname + s.Next().AttrOr("href", "NotFound")
				return
			}
		},
	)
	return
}
