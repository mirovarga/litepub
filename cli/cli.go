package cli

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

const (
	postsDir     = "posts"
	templatesDir = "templates"
	outputDir    = "www"
)

var log quietableLog

func Run() int {
	arguments, _ := docopt.ParseArgs(usage, nil, "LitePub 0.5.4")

	log = quietableLog{arguments["--quiet"].(int) == 1}

	if _, ok := arguments["<dir>"].(string); !ok {
		arguments["<dir>"] = "."
	}

	if arguments["create"].(bool) {
		return create(arguments)
	} else if arguments["build"].(bool) {
		return build(arguments)
	} else if arguments["serve"].(bool) {
		return serve(arguments)
	}

	return 0
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
