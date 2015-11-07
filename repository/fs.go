package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "mirovarga.com/litepub"
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
	return fmt.Errorf("Not implemented")
}

func (r FSBlogRepository) FindByID(id string) (Blog, error) {
	if _, err := os.Stat(filepath.Join(r.dir, id)); err != nil {
		return Blog{}, fmt.Errorf("Blog not found: %s", id)
	}

	postFiles, err := ioutil.ReadDir(filepath.Join(r.dir, postsDir))
	if err != nil {
		return Blog{}, fmt.Errorf("Failed to read posts: %s", err)
	}

	blog := Blog{ID: id}
	for _, postFile := range postFiles {
		post, err := r.parsePost(postFile.Name())
		if err != nil {
			return Blog{}, err
		}
		blog.Posts = append(blog.Posts, post)
	}

	return blog, nil
}

func (r FSBlogRepository) parsePost(fileName string) (Post, error) {
	markdown, err := ioutil.ReadFile(filepath.Join(postsDir, fileName))
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
