package main

import (
	"fmt"

	. "mirovarga.com/litepub"
	"mirovarga.com/litepub/adapters"
)

// TODO make dirs overridable via command line arguments (here or at create command) (?)
// TODO multiple templates (specify as command line arguments)
// TODO -z, --zip [<path>]  Zip the www directory to the archive defined by path [default: www.zip]
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
