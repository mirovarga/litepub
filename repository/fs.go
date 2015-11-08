package repository

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

func NewFSBlogRepository(dir string) BlogRepository {
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0700)
	}
	return FSBlogRepository{dir}
}

const postsDir = "posts"

type FSBlogRepository struct {
	dir string
}

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
