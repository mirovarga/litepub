package main

import (
	"fmt"

	. "mirovarga.com/litepub"
	"mirovarga.com/litepub/adapters"
)

// TODO multiple templates (specify as command line arguments)
// TODO -o, --output <dir>  Generate the blog to the specified directory [default: www]
// TODO -z, --zip [<file>]  Zip the <dir> directory to an archive [default: <dir>.zip]
func build(arguments map[string]interface{}) {
	blogRepo := adapters.NewFSBlogRepository(".")
	readers := NewReaders(blogRepo)

	gen, err := adapters.NewStaticBlogGeneratorWithProgress(
		"", templatesDir, outputDir, func(fileName string) {
			fmt.Printf("Generating: %s\n", fileName)
		}, readers)
	if err != nil {
		fmt.Printf("Failed to create static generator: %s\n", err)
		return
	}

	err = gen.Generate()
	if err != nil {
		fmt.Printf("Failed to generate blog: %s\n", err)
	}
}
