package main

import "github.com/docopt/docopt-go"

// TODO server --watch, --port
// TODO create <name=litepub-blog> --empty
// TODO sample blog with templates and posts dirs in releases
// TODO linux, win & mac releases: http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub, 0.1.0", false)

	if arguments["build"].(bool) {
		build(arguments)
	}
}
