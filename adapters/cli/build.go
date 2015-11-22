package main

import (
	"fmt"
	"path/filepath"

	"mirovarga.com/litepub/adapters"
	"mirovarga.com/litepub/application"
)

// TODO -o, --output <dir>  Generate the blog to the specified directory [default: www]
func build(arguments map[string]interface{}) {
	dir := arguments["<dir>"].(string)

	blogRepo := adapters.NewFSBlogRepository(repoDir(dir))
	readers := application.NewReaders(blogRepo)

	gen, err := adapters.NewStaticBlogGeneratorWithProgress(blogID(dir),
		filepath.Join(dir, templatesDir), filepath.Join(dir, outputDir),
		printProgress, readers)
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
