package main

import (
	"path/filepath"

	"mirovarga.com/litepub/adapters"
	"mirovarga.com/litepub/application"
)

func build(arguments map[string]interface{}) {
	dir := arguments["<dir>"].(string)

	blogRepo := adapters.NewFSBlogRepository(repoDir(dir))
	readers := application.NewReaders(blogRepo)

	gen, err := adapters.NewStaticBlogGeneratorWithProgress(blogID(dir),
		filepath.Join(dir, templatesDir), filepath.Join(dir, outputDir),
		printProgress, readers)
	if err != nil {
		log.Fatalf("Failed to create generator: %s\n", err)
		return
	}

	err = gen.Generate()
	if err != nil {
		log.Fatalf("Failed to generate blog: %s\n", err)
	}
}

func printProgress(path string) {
	log.Printf("Generating: %s\n", path)
}
