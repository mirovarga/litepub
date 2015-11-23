package main

import (
	"path/filepath"

	"github.com/docopt/docopt-go"
)

const (
	postsDir     = "posts"
	templatesDir = "templates"
	outputDir    = "www"
)

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub, 0.4.0", false)

	if _, ok := arguments["<dir>"].(string); !ok {
		arguments["<dir>"] = "."
	}

	if arguments["create"].(bool) {
		create(arguments)
	} else if arguments["build"].(bool) {
		build(arguments)
	} else if arguments["serve"].(bool) {
		serve(arguments)
	}
}

func repoDir(dir string) string {
	dirs := filepath.SplitList(dir)
	return filepath.Join(dirs[:len(dirs)-1]...)
}

func blogID(dir string) string {
	dirs := filepath.SplitList(dir)
	return dirs[len(dirs)-1]
}
