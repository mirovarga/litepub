package main

import "github.com/docopt/docopt-go"

const (
	postsDir     = "posts"
	templatesDir = "templates"
	outputDir    = "www"
)

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub, 0.3.0", false)

	if arguments["create"].(bool) {
		create(arguments)
	} else if arguments["build"].(bool) {
		build(arguments)
	} else if arguments["serve"].(bool) {
		serve(arguments)
	}
}
