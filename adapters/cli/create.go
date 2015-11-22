package main

//go:generate go-bindata -prefix sample/ -o sample.go sample/...

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"mirovarga.com/litepub/adapters"
	"mirovarga.com/litepub/application"
)

const defaultName = "litepub-blog"

var templates = []string{"layout.tmpl", "index.tmpl", "post.tmpl", "tag.tmpl"}

func create(arguments map[string]interface{}) {
	blogRepo := adapters.NewFSBlogRepository(".")
	authors := application.NewAuthors(blogRepo)

	name, ok := arguments["<name>"].(string)
	if !ok {
		name = defaultName
	}

	err := authors.CreateBlog(name)
	if err != nil {
		fmt.Printf("Failed to create blog: %s\n", err)
		return
	}

	if arguments["--blank"].(int) == 1 {
		dir := filepath.Join(name, templatesDir)
		os.MkdirAll(dir, 0700)

		for _, template := range templates {
			ioutil.WriteFile(filepath.Join(dir, template), nil, 0600)
		}
	} else {
		RestoreAssets(name, "templates")
		RestoreAssets(name, "posts")
	}
}
