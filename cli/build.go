package cli

import (
	"path/filepath"

	"mirovarga.com/litepub/lib"
)

func build(arguments map[string]interface{}) int {
	dir := arguments["<dir>"].(string)

	blog, err := lib.NewMarkdownBlog(dir).Read()
	if err != nil {
		log.Errorf("Failed to read blog: %s\n", err)
		return -1
	}

	gen, err := lib.NewStaticBlogGenerator(blog, filepath.Join(dir, templatesDir),
		filepath.Join(dir, outputDir), printProgress)
	if err != nil {
		log.Errorf("Failed to create generator: %s\n", err)
		return -1
	}

	err = gen.Generate()
	if err != nil {
		log.Errorf("Failed to generate blog: %s\n", err)
		return -1
	}

	return 0
}

func printProgress(path string) {
	log.Infof("Generating: %s\n", path)
}
