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
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Infoln("finding all URLs we need to crawl")
	cs.FoundURLs = c.findURLsToDepth()
	logrus.Infof("found %d URLs crawling for words", len(cs.FoundURLs))

}
