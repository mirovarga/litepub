package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/gosimple/slug"
	"github.com/russross/blackfriday"
)

const usage = `
LitePub - a lightweight static blog generator

Usage:
  litepub generate

Options:
  -h, --help      Show this screen
  -v, --version   Show version
`

// TODO server command
// XXX split to parser, generator/template, cli, ...
func main() {
	arguments, _ := docopt.Parse(usage, nil, true, "LitePub, 0.1.0", false)

	if arguments["generate"].(bool) {
		generate()
	}
}

// TODO make overridable via command line arguments
const (
	templatesDir = "templates"
	postsDir     = "posts"
	wwwDir       = "www"
)

var tmplFuncs = template.FuncMap{
	"html": func(markdown string) template.HTML {
		html := blackfriday.MarkdownCommon([]byte(markdown))
		return template.HTML(html)
	},
	"summary": func(content string) string {
		lines := strings.Split(content, "\n\n")
		for _, line := range lines {
			if !strings.HasPrefix(line, "#") {
				return line
			}
		}
		return content
	},
	"even": func(index int) bool {
		return index%2 == 0
	},
	"inc": func(index int) int {
		return index + 1
	},
}

// TODO tags
// TODO multiple templates (specify as command line argument)
func generate() {
	os.RemoveAll(wwwDir)

	// FIXME use os-independent way
	err := exec.Command("cp", "-r", templatesDir, wwwDir).Run()
	if err != nil {
		log.Fatalf("Error copying templates: %s", err)
	}

	os.Remove(filepath.Join(wwwDir, "layout.tmpl"))
	os.Remove(filepath.Join(wwwDir, "index.tmpl"))
	os.Remove(filepath.Join(wwwDir, "post.tmpl"))

	postFiles, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Fatalf("Error reading posts: %s", err)
	}

	var posts []post
	for _, postFile := range postFiles {
		posts = append(posts, generatePost(postFile.Name()))
	}

	generateIndex(posts)
}

func generateIndex(posts []post) {
	log.Println("Generating home page")

	sort.Sort(sortedPosts(posts))

	tmpl, err := template.New("layout.tmpl").Funcs(tmplFuncs).ParseFiles(
		filepath.Join(templatesDir, "layout.tmpl"),
		filepath.Join(templatesDir, "index.tmpl"))

	f, err := os.OpenFile(filepath.Join(wwwDir, "index.html"), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Error generating home page: %s", err)
	}
	defer f.Close()

	tmpl.Execute(f, posts)
}

func generatePost(fileName string) post {
	log.Printf("Generating post: %s", fileName)

	markdown, err := ioutil.ReadFile(filepath.Join(postsDir, fileName))
	if err != nil {
		log.Fatalf("Error generating post: %s", err)
	}

	post := parsePost(string(markdown))

	tmpl, err := template.New("layout.tmpl").Funcs(tmplFuncs).ParseFiles(
		filepath.Join(templatesDir, "layout.tmpl"),
		filepath.Join(templatesDir, "post.tmpl"))

	f, err := os.OpenFile(filepath.Join(wwwDir, post.Slug+".html"), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Error generating post: %s", err)
	}
	defer f.Close()

	tmpl.Execute(f, post)

	return post
}

func parsePost(markdown string) post {
	lines := strings.Split(markdown, "\n\n")

	title := strings.Replace(lines[0], "#", "", -1)
	written, _ := time.Parse("*Jan 2, 2006*", lines[1])
	content := strings.Join(lines[2:], "\n\n")

	return post{
		Title:   title,
		Written: written,
		Content: content,
		Slug:    slug.Make(title),
	}
}

type post struct {
	// Plain text
	Title   string
	Written time.Time

	// Markdown
	Content string

	// Generated from the title
	Slug string
}

type sortedPosts []post

func (s sortedPosts) Len() int {
	return len(s)
}

func (s sortedPosts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedPosts) Less(i, j int) bool {
	return s[j].Written.Before(s[i].Written)
}
