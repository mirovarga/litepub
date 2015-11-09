package adapters

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "mirovarga.com/litepub"

	"github.com/gosimple/slug"
)

const postsDir = "posts"

// FSBlogRepository stores Blogs on the file system.
//
// Blogs are stored as subdirectories of the main directory. The subdirectories
// are named after the IDs of the Blogs they contain. Posts are stored as Markdown
// files in the 'posts' subdirectory of their Blog directory.
//
// So the structure looks like this:
//
// main directory, fe. blogs
// 	 blog1
//     posts
//       post1.md
//       post2.md
//       ...
//   blog2
//     posts
//       post1.md
//       post2.md
//       ...
//   ...
type FSBlogRepository struct {
	dir string
}

// Store adds or updates a Blog in the repository. It creates all necessary
// directories if needed.
func (r FSBlogRepository) Store(blog Blog) error {
	blogDir := filepath.Join(r.dir, blog.ID)
	os.MkdirAll(blogDir, 0700)
	os.Mkdir(filepath.Join(blogDir, postsDir), 0700)

	for _, post := range blog.Posts {
		err := r.writePost(blog.ID, post)
		if err != nil {
			return fmt.Errorf("Failed to store post: %s", err)
		}
	}

	return nil
}

// FindByID gets a Blog with the ID from the repository.
//
// If the Blog isn't in the repository it returns an error.
func (r FSBlogRepository) FindByID(id string) (Blog, error) {
	blogDir := filepath.Join(r.dir, id)

	if _, err := os.Stat(blogDir); err != nil {
		return Blog{}, fmt.Errorf("Blog not found: %s", id)
	}

	postFiles, err := ioutil.ReadDir(filepath.Join(blogDir, postsDir))
	if err != nil {
		return Blog{}, fmt.Errorf("Failed to read posts: %s", err)
	}

	blog := Blog{ID: id}
	for _, postFile := range postFiles {
		post, err := r.readPost(id, postFile.Name())
		if err != nil {
			return Blog{}, err
		}
		blog.Posts = append(blog.Posts, post)
	}

	return blog, nil
}

func (r FSBlogRepository) writePost(id string, post Post) error {
	data := fmt.Sprintf("# %s\n\n*%s*\n\n%s\n",
		post.Title, post.Written.Format("Jan 2, 2006"), post.Content)
	return ioutil.WriteFile(filepath.Join(r.dir, id, postsDir,
		slug.Make(post.Title)+".md"), []byte(data), 0600)
}

func (r FSBlogRepository) readPost(id, fileName string) (Post, error) {
	markdown, err := ioutil.ReadFile(filepath.Join(r.dir, id, postsDir, fileName))
	if err != nil {
		return Post{}, fmt.Errorf("Failed to read post: %s", err)
	}

	lines := strings.Split(string(markdown), "\n\n")

	title := strings.TrimSpace(strings.Replace(lines[0], "#", "", -1))

	written, err := time.Parse("*Jan 2, 2006*", lines[1])
	if err != nil {
		return Post{}, fmt.Errorf("Failed to parse post date: %s", err)
	}

	content := strings.Join(lines[2:], "\n\n")

	return Post{title, content, written}, nil
}

// NewFSBlogRepository creates a FSBlogRepository in the provided directory. If
// the directory doesn't exist it creates it.
func NewFSBlogRepository(dir string) BlogRepository {
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0700)
	}
	return FSBlogRepository{dir}
}
