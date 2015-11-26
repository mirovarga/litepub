package main

import (
	"fmt"
	"path/filepath"

	"github.com/docopt/docopt-go"
)

const (
	postsDir     = "posts"
	templatesDir = "templates"
	outputDir    = "www"
)

var log quietableLog

func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub 0.5.0", false)

	log = quietableLog{arguments["--quiet"].(int) == 1}

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

type quietableLog struct {
	quiet bool
}

func (l quietableLog) Infof(format string, v ...interface{}) {
	if !l.quiet {
		fmt.Printf(format, v...)
	}
}

func (l quietableLog) Errorf(format string, v ...interface{}) {
	fmt.Printf("ERROR: "+format, v...)
}
