package main

import "github.com/docopt/docopt-go"

// TODO server --watch, --port
// TODO create <name=litepub-blog> --empty
// TODO sample blog with templates and posts dirs in releases

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub, 0.1.0", false)

	if arguments["build"].(bool) {
		build(arguments)
	}
}
