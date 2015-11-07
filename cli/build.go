package main

import (
	"fmt"

	. "mirovarga.com/litepub"
	"mirovarga.com/litepub/generator"
	"mirovarga.com/litepub/repository"
)

const (
	templatesDir = "templates"
	outputDir    = "www"
)

// TODO make dirs overridable via command line arguments (here or at create command) (?)
// TODO multiple templates (specify as command line arguments)
func build(arguments map[string]interface{}) {
	blogRepo := repository.NewFSBlogRepository(".")
	usecases := NewUsecases(blogRepo)

	gen, err := generator.NewStaticBlogGeneratorWithProgress(
		"", templatesDir, outputDir, func(fileName string) {
			fmt.Printf("Generating: %s\n", fileName)
		}, usecases)
	if err != nil {
		fmt.Printf("Failed to create static generator: %s\n", err)
		return
	}

	err = gen.Generate()
	if err != nil {
		fmt.Printf("Failed to generate blog: %s\n", err)
	}
}
