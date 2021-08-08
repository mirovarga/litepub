package cli

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed sample
var assets embed.FS

var templates = []string{"layout.tmpl", "index.tmpl", "post.tmpl", "tag.tmpl"}

func create(arguments map[string]interface{}) int {
	dir := arguments["<dir>"].(string)

	err := os.MkdirAll(dir, 0700)
	if err != nil {
		log.Errorf("Failed to create blog: %s\n", err)
		return -1
	}

	if arguments["--skeleton"].(int) == 1 {
		tmplDir := filepath.Join(dir, templatesDir)
		err := os.MkdirAll(tmplDir, 0700)
		if err != nil {
			log.Errorf("Failed to create blog: %s\n", err)
			return -1
		}

		for _, template := range templates {
			err := os.WriteFile(filepath.Join(tmplDir, template), nil, 0600)
			if err != nil {
				log.Errorf("Failed to create blog: %s\n", err)
				return -1
			}
		}
	} else {
		err := fs.WalkDir(assets, "sample", func(path string, d fs.DirEntry, err error) error {
			if path == "sample" {
				return nil
			}

			trimmedPath := strings.TrimPrefix(path, "sample/")
			if d.IsDir() {
				return os.MkdirAll(filepath.Join(dir, trimmedPath), 0700)
			} else {
				bytes, err := fs.ReadFile(assets, path)
				if err != nil {
					return err
				}
				return os.WriteFile(filepath.Join(dir, trimmedPath), bytes, 0600)
			}
		})
		if err != nil {
			log.Errorf("Failed to create blog: %s\n", err)
			return -1
		}
	}

	log.Infof("Created blog: %s\n", dir)
	return 0
}
