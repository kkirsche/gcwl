package libgcwl

// FlagConfig is used to contain the different flag values that we receive from
// the user on the command line. Including things like help, verbose, depth,
// etc.
type FlagConfig struct {
	Depth          int
	MinWordLength  int
	IncludeEmail   bool
	EmailFile      string
	IncludeMeta    bool
	MetaFile       string
	NoWords        bool
	AllowOffsite   bool
	WriteTo        string
	UserAgent      string
	MetaTempDir    string
	KeepDownloaded bool
	Count          bool
	Verbose        bool
	WorkerCount    int
	SeedURLs       []string
}

// CrawlState is used to track information we need during the crawl
type CrawlState struct {
	FoundURLs       []string
	UnprocessedURLs []string
}
