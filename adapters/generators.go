package adapters

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gosimple/slug"
	"github.com/russross/blackfriday"
	"github.com/termie/go-shutil"

	. "mirovarga.com/litepub"
)

// ProgressFunc is used to monitor progress of generating a Blog. It is called
// before a file generation is started.
type ProgressFunc func(fileName string)

// StaticBlogGenerator generates Blogs to static HTML files.
type StaticBlogGenerator struct {
	id            string
	templatesDir  string
	outputDir     string
	progressFunc  ProgressFunc
	readers       Readers
	indexTemplate *template.Template
	postTemplate  *template.Template
}

// Generate generates a Blog to static HTML files.
func (g StaticBlogGenerator) Generate() error {
	err := g.prepareOutputDir()
	if err != nil {
		return fmt.Errorf("Failed to prepare output directory: %s", err)
	}

	err = g.generateIndex()
	if err != nil {
		return fmt.Errorf("Failed to generate index: %s", err)
	}

	err = g.generatePosts()
	if err != nil {
		return fmt.Errorf("Failed to generate posts: %s", err)
	}

	return nil
}

func (g StaticBlogGenerator) prepareOutputDir() error {
	os.RemoveAll(g.outputDir)

	err := shutil.CopyTree(g.templatesDir, g.outputDir,
		&shutil.CopyTreeOptions{
			Symlinks: true,
			Ignore: func(string, []os.FileInfo) []string {
				return []string{"layout.tmpl", "index.tmpl", "post.tmpl"}
			},
			CopyFunction:           shutil.Copy,
			IgnoreDanglingSymlinks: false,
		})
	if err != nil {
		return err
	}

	return nil
}

func (g StaticBlogGenerator) generateIndex() error {
	blog, err := g.readers.GetBlog(g.id)
	if err != nil {
		return err
	}

	templatePosts := toTemplatePosts(blog.Posts...)
	sort.Sort(sortedTemplatePosts(templatePosts))
	return g.generatePage(g.indexTemplate, "index.html", templatePosts)
}

func (g StaticBlogGenerator) generatePosts() error {
	blog, err := g.readers.GetBlog(g.id)
	if err != nil {
		return err
	}

	for _, post := range blog.Posts {
		templatePost := toTemplatePosts(post)[0]
		err = g.generatePage(g.postTemplate, templatePost.Slug+".html", templatePost)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g StaticBlogGenerator) generatePage(template *template.Template,
	fileName string, data interface{}) error {
	g.progressFunc(fileName)

	pageFile, err := os.OpenFile(filepath.Join(g.outputDir, fileName),
		os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer pageFile.Close()

	return template.Execute(pageFile, data)
}

// NewStaticBlogGenerator creates a StaticBlogGenerator that generates the Blog
// with the ID to static HTML files in the outputDir using templates from
// the templatesDir.
func NewStaticBlogGenerator(id, templatesDir, outputDir string,
	readers Readers) (StaticBlogGenerator, error) {
	return NewStaticBlogGeneratorWithProgress(id, templatesDir, outputDir, nil, readers)
}

// NewStaticBlogGeneratorWithProgress creates a StaticBlogGenerator that
// generates the Blog with the ID to static HTML files in the outputDir using
// templates from the templatesDir. It calls the progressFunc before generating
// each file.
func NewStaticBlogGeneratorWithProgress(id, templatesDir, outputDir string,
	progressFunc ProgressFunc, readers Readers) (StaticBlogGenerator, error) {
	if _, err := os.Stat(templatesDir); err != nil {
		return StaticBlogGenerator{},
			fmt.Errorf("Templates directory not found: %s", templatesDir)
	}

	indexTemplate, err := template.New("layout.tmpl").Funcs(templateFuncs).ParseFiles(
		filepath.Join(templatesDir, "layout.tmpl"),
		filepath.Join(templatesDir, "index.tmpl"))
	if err != nil {
		return StaticBlogGenerator{}, err
	}

	postTemplate, err := template.New("layout.tmpl").Funcs(templateFuncs).ParseFiles(
		filepath.Join(templatesDir, "layout.tmpl"),
		filepath.Join(templatesDir, "post.tmpl"))
	if err != nil {
		return StaticBlogGenerator{}, err
	}

	return StaticBlogGenerator{id, templatesDir, outputDir, progressFunc,
		readers, indexTemplate, postTemplate}, nil
}

type templatePost struct {
	Post
	Slug string
}

func toTemplatePosts(posts ...Post) []templatePost {
	var templatePosts []templatePost
	for _, post := range posts {
		templatePosts = append(templatePosts, templatePost{post, slug.Make(post.Title)})
	}
	return templatePosts
}

type sortedTemplatePosts []templatePost

func (p sortedTemplatePosts) Len() int {
	return len(p)
}

func (p sortedTemplatePosts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p sortedTemplatePosts) Less(i, j int) bool {
	return p[j].Written.Before(p[i].Written)
}

var templateFuncs = template.FuncMap{
	"html":    html,
	"summary": summary,
	"even":    even,
	"inc":     inc,
}

func html(markdown string) template.HTML {
	html := blackfriday.MarkdownCommon([]byte(markdown))
	return template.HTML(html)
}

func summary(content string) string {
	lines := strings.Split(content, "\n\n")
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			return line
		}
	}
	return content
}

func even(integer int) bool {
	return integer%2 == 0
}

func inc(integer int) int {
	return integer + 1
}
