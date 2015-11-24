package main

//go:generate go-bindata -prefix sample/ -o sample.go sample/...

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"mirovarga.com/litepub/adapters"
	"mirovarga.com/litepub/application"
)

var templates = []string{"layout.tmpl", "index.tmpl", "post.tmpl", "tag.tmpl"}

func create(arguments map[string]interface{}) {
	dir := arguments["<dir>"].(string)

	blogRepo := adapters.NewFSBlogRepository(repoDir(dir))
	authors := application.NewAuthors(blogRepo)

	err := authors.CreateBlog(blogID(dir))
	if err != nil {
		log.Fatalf("Failed to create blog: %s\n", err)
		return
	}

	if arguments["--skeleton"].(int) == 1 {
		tmplDir := filepath.Join(dir, templatesDir)
		os.MkdirAll(tmplDir, 0700)

		for _, template := range templates {
			ioutil.WriteFile(filepath.Join(tmplDir, template), nil, 0600)
		}
	} else {
		RestoreAssets(dir, "templates")
		RestoreAssets(dir, "posts")
	}
	log.Printf("Created blog: %s\n", dir)
}
