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
	// d, v
	RootCmd.Flags().IntVarP(&config.Depth, "depth", "d", 2, "depth to spider the seed URLs to")
	RootCmd.Flags().BoolVarP(&config.Verbose, "verbose", "v", false, "enable verbose logging")
}
