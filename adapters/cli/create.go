package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"mirovarga.com/litepub/adapters"
	"mirovarga.com/litepub/application"
)

const (
	layoutTemplate = `<!DOCTYPE html>

<html>

  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>LitePub Blog: {{template "title" .}}</title>

    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Raleway:400,600,300">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/3.0.3/normalize.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css">
  </head>

  <body>
    <div class="container">
      <header>
        <div class="row">
          <div class="twelve columns">
            <h1><a href="/">LitePub Blog</a></h1>
            <hr>
          </div>
        </div>
      </header>

      {{template "content" .}}

      <footer>
        <div class="row">
          <div class="twelve columns">
            <hr>
            &copy; 2015 LitePub
          </div>
        </div>
      </footer>
    </div>
  </body>

</html>`
	indexTemplate = `{{define "title"}}
  Home
{{end}}

{{define "content"}}
  {{range .}}
    <div class="row">
      <div class="twelve columns">
        <h4><a href="/{{.Title | slug}}.html">{{.Title}}</a></h4>
        {{.Content | summary | html}}
      </div>
    </div>
  {{end}}
{{end}}`
	postTemplate = `{{define "title"}}
  {{.Title}}
{{end}}

{{define "content"}}
  <div class="row">
    <div class="twelve columns">
      <h1>{{.Title}}</h1>
      <p>
        <em>{{.Written.Format "Jan 2, 2006"}}</em>
      </p>
	  {{if (len .Tags) ne 0}}
		<p>
		  {{range .Tags}}
			<em><a href="/tags/{{. | slug}}.html">{{.}}</a></em>&nbsp;
		  {{end}}
		</p>
	  {{end}}
      {{.Content | html}}
    </div>
  </div>
{{end}}`
	tagTemplate = `{{define "title"}}
  Posts tagged {{.Name}}
{{end}}

{{define "content"}}
  <div class="row">
    <div class="twelve columns">
      <h1>{{template "title" .}}</h1>
	</div>
  </div>

  {{range .Posts}}
    <div class="row">
      <div class="twelve columns">
        <h4><a href="/{{.Title | slug}}.html">{{.Title}}</a></h4>
        {{.Content | summary | html}}
      </div>
    </div>
  {{end}}
{{end}}`
)

const defaultName = "litepub-blog"

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

	os.Mkdir(filepath.Join(name, templatesDir), 0700)

	dir := filepath.Join(name, templatesDir)
	if arguments["--blank"].(int) == 1 {
		writeTemplate(filepath.Join(dir, "layout.tmpl"), "")
		writeTemplate(filepath.Join(dir, "index.tmpl"), "")
		writeTemplate(filepath.Join(dir, "post.tmpl"), "")
		writeTemplate(filepath.Join(dir, "tag.tmpl"), "")
	} else {
		writeTemplate(filepath.Join(dir, "layout.tmpl"), layoutTemplate)
		writeTemplate(filepath.Join(dir, "index.tmpl"), indexTemplate)
		writeTemplate(filepath.Join(dir, "post.tmpl"), postTemplate)
		writeTemplate(filepath.Join(dir, "tag.tmpl"), tagTemplate)

		err = authors.CreatePost(name, "Welcome to LitePub!",
			"LitePub is a lightweight static blog generator written in Go.",
			time.Now(), "LitePub")
		if err != nil {
			fmt.Printf("Failed to create post: %s\n", err)
		}
	}
}

func writeTemplate(filePath, content string) {
	err := ioutil.WriteFile(filePath, []byte(content), 0600)
	if err != nil {
		fmt.Printf("Failed to write template: %s\n", err)
	}
}
