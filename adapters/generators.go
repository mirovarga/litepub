package adapters

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gosimple/slug"
	"github.com/russross/blackfriday"
	"github.com/termie/go-shutil"

	"mirovarga.com/litepub/application"
	"mirovarga.com/litepub/domain"
)

// ProgressFunc is used to monitor progress of generating a Blog. It is called
// before a file generation is started.
type ProgressFunc func(path string)

// StaticBlogGenerator generates Blogs to static HTML files.
type StaticBlogGenerator struct {
	id            string
	templatesDir  string
	outputDir     string
	progressFunc  ProgressFunc
	readers       application.Readers
	indexTemplate *template.Template
	postTemplate  *template.Template
	tagTemplate   *template.Template
	posts         []domain.Post
	postsByTag    map[string][]domain.Post
}

// Generate generates a Blog to static HTML files.
func (g StaticBlogGenerator) Generate() error {
	err := g.prepareOutputDir()
	if err != nil {
		return fmt.Errorf("Failed to prepare output directory: %s", err)
	}

	err = g.readPosts()
	if err != nil {
		return fmt.Errorf("Failed to read posts: %s", err)
	}

	err = g.generateIndex()
	if err != nil {
		return fmt.Errorf("Failed to generate index: %s", err)
	}

	err = g.generateTags()
	if err != nil {
		return fmt.Errorf("Failed to generate tags: %s", err)
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
				return []string{"layout.tmpl", "index.tmpl", "post.tmpl", "tag.tmpl"}
			},
			CopyFunction:           shutil.Copy,
			IgnoreDanglingSymlinks: false,
		})
	if err != nil {
		return err
	}

	os.Mkdir(filepath.Join(g.outputDir, "tags"), 0700)

	return nil
}

func (g *StaticBlogGenerator) readPosts() error {
	blog, err := g.readers.GetBlog(g.id)
	if err != nil {
		return err
	}

	g.posts = blog.PostsByDate(false, false)

	for _, tag := range blog.Tags(false) {
		g.postsByTag[tag] = blog.PostsByDate(false, false, tag)
	}

	return nil
}

func (g StaticBlogGenerator) generateIndex() error {
	return g.generatePage(g.indexTemplate, "index.html", g.posts)
}

func (g StaticBlogGenerator) generatePosts() error {
	for _, post := range g.posts {
		err := g.generatePage(g.postTemplate, slug.Make(post.Title)+".html", post)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g StaticBlogGenerator) generateTags() error {
	for tag, posts := range g.postsByTag {
		err := g.generatePage(g.tagTemplate,
			filepath.Join("tags", slug.Make(tag)+".html"), struct {
				Name  string
				Posts []domain.Post
			}{tag, posts})
		if err != nil {
			return err
		}
	}

	return nil
}

func (g StaticBlogGenerator) generatePage(template *template.Template,
	path string, data interface{}) error {
	g.progressFunc(path)

	pageFile, err := os.OpenFile(filepath.Join(g.outputDir, path),
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
	readers application.Readers) (StaticBlogGenerator, error) {
	return NewStaticBlogGeneratorWithProgress(id, templatesDir, outputDir, nil, readers)
}

// NewStaticBlogGeneratorWithProgress creates a StaticBlogGenerator that
// generates the Blog with the ID to static HTML files in the outputDir using
// templates from the templatesDir. It calls the progressFunc before generating
// each file.
func NewStaticBlogGeneratorWithProgress(id, templatesDir, outputDir string,
	progressFunc ProgressFunc, readers application.Readers) (StaticBlogGenerator, error) {
	if _, err := os.Stat(templatesDir); err != nil {
		return StaticBlogGenerator{},
			fmt.Errorf("Templates directory not found: %s", templatesDir)
	}

	indexTemplate, err := createTemplate(templatesDir, "index.tmpl")
	if err != nil {
		return StaticBlogGenerator{}, err
	}

	postTemplate, err := createTemplate(templatesDir, "post.tmpl")
	if err != nil {
		return StaticBlogGenerator{}, err
	}

	tagTemplate, err := createTemplate(templatesDir, "tag.tmpl")
	if err != nil {
		return StaticBlogGenerator{}, err
	}

	return StaticBlogGenerator{id, templatesDir, outputDir, progressFunc,
		readers, indexTemplate, postTemplate, tagTemplate, []domain.Post{},
		make(map[string][]domain.Post)}, nil
}

func createTemplate(dir, name string) (*template.Template, error) {
	return template.New("layout.tmpl").Funcs(templateFuncs).ParseFiles(
		filepath.Join(dir, "layout.tmpl"),
		filepath.Join(dir, name))
}

var templateFuncs = template.FuncMap{
	"html":    html,
	"summary": summary,
	"even":    even,
	"inc":     inc,
	"slug":    slugify,
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

func slugify(str string) string {
	return slug.Make(str)
}
