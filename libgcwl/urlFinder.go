package libgcwl

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

// findURLsToDepth is used to find all URLs that we will need to visit up to a
// specific depth
func (c *FlagConfig) findURLsToDepth() []string {
	urls := c.SeedURLs
	results := c.SeedURLs
	for i := 0; i <= c.Depth; i++ {
		logrus.WithField("remaining-depth", c.Depth-i).Debugln("crawling...")
		urls = crawl(urls)
		results = append(results, urls...)
	}

	return results
}

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

func crawl(urls []string) []string {
	var foundURLs []string

	// Channels
	chUrls := make(chan string)
	chErrs := make(chan error)
	chFinished := make(chan bool)

	// Kick off the crawl process (concurrently)
	for _, url := range urls {
		go execCrawl(url, chUrls, chErrs, chFinished)
	}

	// Subscribe to both channels
	for c := 0; c < len(urls); {
		select {
		case url := <-chUrls:
			foundURLs = append(foundURLs, url)
		case err := <-chErrs:
			logrus.WithError(err).Errorln("error crawling to url")
		case <-chFinished:
			c++
			logrus.WithField("remaining", len(urls)-c).Debugln("remaining URLs in level")
		}
	}

	close(chUrls)

	return foundURLs
}

func execCrawl(url string, ch chan string, errs chan error, chFinished chan bool) {
	resp, err := http.Get(url)

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	if err != nil {
		errs <- err
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}

			// Make sure the url begines in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}
		}
	}
}
