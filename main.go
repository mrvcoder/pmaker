package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

var (
	help         = flag.Bool("h", false, "help menu")
	params_count = flag.Int("count", 15, "Count of params per req to send (default: 15)")
	params_value = flag.String("value", "vcoder", "vlaue to set for all params (default: vcoder)")
)

func main() {
	flag.Parse()
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)

	if *help == true {
		fmt.Println(`Param Maker - Make Parameters From List Of Words :D
-----------------------------------------------------------------------------
cat wordlists | pmaker [options]
--count						Count of params per req to send (default: 15)
--value						vlaue to set for all params (default: vcoder)
Follow me at github => https://github.com/mrvcoder
Follow me at twitter => https://twitter.com/VC0D3R
`)
		os.Exit(0)
	}

	Stdin := ReadStdin()

	params := strings.Split(Stdin, "\n")
	chunks := ChunkStringSlice(params, *params_count)
	// Build the query strings for each chunk
	for i, chunk := range chunks {
		d := CreateParams(chunk)
		out := fmt.Sprintf("Part %d:\n\n%s\n\n", i+1, d)
		fmt.Println(out)
	}
}

func CreateParams(params []string) string {
	// Build the query string
	var queryParams []string
	for _, word := range params {
		queryParams = append(queryParams, fmt.Sprintf("%s=%s", word, *params_value))
	}
	queryString := strings.Join(queryParams, "&")
	return queryString
}
