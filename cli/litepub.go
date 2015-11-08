package main

import "github.com/docopt/docopt-go"

const (
	templatesDir = "templates"
	outputDir    = "www"
)

// TODO server --watch, --port
func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub, 0.1.0", false)

	if arguments["create"].(bool) {
		create(arguments)
	} else if arguments["build"].(bool) {
		build(arguments)
	}
}
