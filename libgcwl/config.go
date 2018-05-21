package libgcwl

// FlagConfig is used to contain the different flag values that we receive from
// the user on the command line. Including things like help, verbose, depth,
// etc.
type FlagConfig struct {
	SeedURLs []string `json:"seed_urls",xml:"seed_urls"`
	Depth    int      `json:"depth",xml:"depth"`
	Verbose  bool     `json:"verbose",xml:"verbose"`
}

// CrawlState is used to track information we need during the crawl
type CrawlState struct {
	FoundURLs       map[string]struct{}
	UnprocessedURLs map[string]struct{}
	FoundWords      map[string]struct{}
}
