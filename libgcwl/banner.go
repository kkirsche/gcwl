package libgcwl

import "fmt"

// PrintBanner outputs the welcome banner to users using the tool
func PrintBanner() {
	banner := fmt.Sprintf(
		"GCWL %s\nKevin Kirsche (d3c3pt10n@deceiveyour.team)\nhttps://deceiveyour.team",
		version)
	fmt.Println(banner)
}
