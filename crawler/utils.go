package crawler

import (
	"crypto/tls"
	"net/http"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Novel represent a novel
type Novel struct {
	Name     string
	IndexURL string
}

// ICrawler - crawl novels
type ICrawler interface {
	LatestChapter(Novel) (string, string)
	AddNovels([]Novel)
	GetNovels() []Novel
}

func gbkToUTF8(gbk string) string {
	gbkToUTF8 := simplifiedchinese.GBK.NewDecoder()
	utf8, _, _ := transform.String(gbkToUTF8, gbk)
	return utf8
}

func getClient() *http.Client {
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return client
}
