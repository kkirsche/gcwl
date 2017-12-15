// Copyright © 2017 Kevin Kirsche <d3c3pt10n@deceiveyour.team>
//
// Licensed under the GNU GPLv3 license;
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/gpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/kkirsche/gcwl/libgcwl"
	"github.com/spf13/cobra"
)

var config libgcwl.FlagConfig

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gcwl",
	Short: "Go Custom Wordlist Generator",
	Long: `gcwl is a Go-based custom wordlist generator based on CeWL. The goal
of this project is to improve on the speed of CeWL through concurrency and by
removing the overhead associated with the Ruby programming language.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		libgcwl.PrintBanner()
		config.SeedURLs = args

		config.RunCrawler()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Shorthands used:
	// a, c, d, e, h, i, k, m, n, o, r, u, w, v
	RootCmd.Flags().IntVarP(&config.Depth, "depth", "d", 2, "depth to spider to")
	RootCmd.Flags().IntVarP(&config.MinWordLength, "min_word_length", "m", 3, "minimum word length")

	RootCmd.Flags().BoolVarP(&config.IncludeEmail, "email", "e", false, "include any email addresses found during the spider")
	RootCmd.Flags().StringVarP(&config.EmailFile, "email_file", "f", "", "Optional output file if email address collection is enabled")

	RootCmd.Flags().BoolVarP(&config.IncludeMeta, "meta", "a", false, "include any meta data found during the spider")
	RootCmd.Flags().StringVarP(&config.MetaFile, "meta_file", "i", "", "Optional output file if meta information collection is enabled")
	RootCmd.Flags().StringVarP(&config.MetaTempDir, "meta-temp-dir", "r", "/tmp", "the temporary directory used when parsing files")

	RootCmd.Flags().BoolVarP(&config.NoWords, "no-words", "n", false, "don't output the wordlist")
	RootCmd.Flags().BoolVarP(&config.AllowOffsite, "offsite", "o", false, "let the spider visit other sites")

	RootCmd.Flags().StringVarP(&config.WriteTo, "write", "w", "", "Write the words to the file")
	RootCmd.Flags().StringVarP(&config.UserAgent, "ua", "u", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36", "The user agent to send")

	RootCmd.Flags().BoolVarP(&config.KeepDownloaded, "keep", "k", false, "keep the documents that are downloaded")
	RootCmd.Flags().BoolVarP(&config.Count, "count", "c", false, "show the count for each of the words found")

	RootCmd.Flags().BoolVarP(&config.Verbose, "verbose", "v", false, "Enable verbose mode")
	RootCmd.Flags().IntVarP(&config.WorkerCount, "worker-threads", "t", 4, "The numer of workers to run in parallel")
}
