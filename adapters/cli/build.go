package main

import (
	"fmt"

	"mirovarga.com/litepub/adapters"
	"mirovarga.com/litepub/application"
)

// TODO -o, --output <dir>  Generate the blog to the specified directory [default: www]
// TODO -z, --zip [<file>]  Zip the <dir> directory to an archive [default: <dir>.zip]
func build(arguments map[string]interface{}) {
	blogRepo := adapters.NewFSBlogRepository(".")
	readers := application.NewReaders(blogRepo)

	gen, err := adapters.NewStaticBlogGeneratorWithProgress("", templatesDir,
		outputDir, printProgress, readers)
	if err != nil {
		fmt.Printf("Failed to create generator: %s\n", err)
		return
	}

	err = gen.Generate()
	if err != nil {
		fmt.Printf("Failed to generate blog: %s\n", err)
	}
}

func printProgress(path string) {
	fmt.Printf("Generating: %s\n", path)
}
