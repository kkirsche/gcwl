package libgcwl

import (
	"github.com/sirupsen/logrus"
)

var (
	cs  CrawlState
	err error
)

// RunCrawler begins the crawling tool
func (c *FlagConfig) RunCrawler() {
	if c.Verbose {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Infoln("finding all URLs we need to crawl")
	cs.FoundURLs = c.findURLsToDepth()
	logrus.Infof("found %d URLs that will require crawling", len(cs.FoundURLs))
}
