package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	postsDir = "posts"
	draftDir = "draft"
)

// MarkdownBlog represents a Blog stored as Markdown files in a directory.
//
// Posts are stored as Markdown files in the posts subdirectory of the Blog
// directory. Draft Posts (ones with Draft set to true) are stored in the draft
// subdirectory of the posts directory.
//
// So the structure looks like this:
//
// blog/
//   posts/
//     draft/
//       draft1.md
//       ...
//     post1.md
//     post2.md
//     ...
//
// Markdown files have the following format:
//
// # Title
//
// *Jan 2, 2009*
//
// *tag1, tag2, ...*
//
// Content
type MarkdownBlog struct {
	dir string
}

// NewMarkdownBlog creates a MarkdownBlog in the provided directory.
//
// If the directory doesn't exist it creates it.
func NewMarkdownBlog(dir string) MarkdownBlog {
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(filepath.Join(dir, postsDir, draftDir), 0700)
	}
	return MarkdownBlog{dir}
}

// Read creates a Blog from the Markdown files.
//
// If the directory doesn't exist it returns an error.
func (b MarkdownBlog) Read() (Blog, error) {
	if _, err := os.Stat(b.dir); err != nil {
		return Blog{}, fmt.Errorf("blog not found: %s", b)
	}

	postsPath := filepath.Join(b.dir, postsDir)
	posts, err := readPosts(postsPath)
	if err != nil {
		return Blog{}, err
	}

	draftsPath := filepath.Join(postsPath, draftDir)
	drafts, err := readPosts(draftsPath)
	if err != nil {
		return Blog{}, err
	}

	blog := posts
	for _, draft := range drafts {
		draft.Draft = true
		blog = append(blog, draft)
	}

	return blog, nil
}

func readPosts(dir string) ([]Post, error) {
	postFiles, err := os.ReadDir(dir)
	if err != nil {
		return []Post{}, fmt.Errorf("failed to read posts: %s", err)
	}

	var posts []Post
	for _, postFile := range postFiles {
		if postFile.IsDir() || strings.HasPrefix(postFile.Name(), ".") {
			continue
		}

		post, err := readPost(filepath.Join(dir, postFile.Name()))
		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func readPost(path string) (Post, error) {
	markdown, err := os.ReadFile(path)
	if err != nil {
		return Post{}, fmt.Errorf("failed to read post: %s", err)
	}

	return markdownToPost(string(markdown))
}

func markdownToPost(markdown string) (Post, error) {
	md := strings.ReplaceAll(markdown, "\r\n", "\n")

	paras := strings.Split(md, "\n\n")
	if len(paras) < 3 {
		return Post{}, fmt.Errorf("title, date or content is missing")
	}

	title := strings.TrimSpace(strings.Replace(paras[0], "#", "", -1))

	written, err := time.Parse("*Jan 2, 2006*", paras[1])
	if err != nil {
		return Post{}, fmt.Errorf("failed to parse date: %s", err)
	}

	var tags []string
	if strings.HasPrefix(paras[2], "*") && !strings.Contains(paras[2], "\n") {
		tags = strings.Split(paras[2], ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(strings.Replace(tag, "*", "", -1))
		}
	}

	var content string
	if len(tags) == 0 {
		content = strings.Join(paras[2:], "\n\n")
	} else {
		content = strings.Join(paras[3:], "\n\n")
	}

	return Post{title, content, written, tags, false}, nil
}
