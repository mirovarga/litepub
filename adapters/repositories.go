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

const (
	postsDir = "posts"
	draftDir = "draft"
)

// FSBlogRepository stores Blogs on the file system.
//
// Blogs are stored as subdirectories of the root directory and are named after
// the IDs of the Blogs they contain. Posts are stored as Markdown files in the
// posts subdirectory of their respective Blog directory. And finally, draft
// posts (ones with Draft set to true) are stored in the draft subdirectory of
// the respective posts directory.
//
// So the structure looks like this:
//
// root directory, for example blogs/
// 	 blog1/
//     posts/
//       draft/
//         draft1.md
//         ...
//       post1.md
//       post2.md
//       ...
//   blog2/
//     posts/
//       draft/
//         draft1.md
//         ...
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
	os.MkdirAll(filepath.Join(r.dir, blog.ID, postsDir, draftDir), 0700)

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

	postsPath := filepath.Join(blogDir, postsDir)
	posts, err := r.readPosts(postsPath)
	if err != nil {
		return Blog{}, err
	}

	draftsPath := filepath.Join(postsPath, draftDir)
	drafts, err := r.readPosts(draftsPath)
	if err != nil {
		return Blog{}, err
	}

	blog := Blog{ID: id}
	blog.Posts = posts
	for _, draft := range drafts {
		draft.Draft = true
		blog.Posts = append(blog.Posts, draft)
	}

	return blog, nil
}

func (r FSBlogRepository) writePost(id string, post Post) error {
	path := filepath.Join(r.dir, id, postsDir)
	if post.Draft {
		path = filepath.Join(path, draftDir)
	}
	path = filepath.Join(path, slug.Make(post.Title)+".md")

	data := fmt.Sprintf("# %s\n\n*%s*\n\n%s\n", post.Title,
		post.Written.Format("Jan 2, 2006"), post.Content)

	return ioutil.WriteFile(path, []byte(data), 0600)
}

func (r FSBlogRepository) readPosts(dir string) ([]Post, error) {
	postFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return []Post{}, fmt.Errorf("Failed to read posts: %s", err)
	}

	var posts []Post
	for _, postFile := range postFiles {
		if postFile.IsDir() {
			continue
		}

		post, err := r.readPost(filepath.Join(dir, postFile.Name()))
		if err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r FSBlogRepository) readPost(path string) (Post, error) {
	markdown, err := ioutil.ReadFile(path)
	if err != nil {
		return Post{}, fmt.Errorf("Failed to read post: %s", err)
	}

	lines := strings.Split(string(markdown), "\n\n")
	if len(lines) < 3 {
		return Post{}, fmt.Errorf("Title, date or content is missing: %s", path)
	}

	title := strings.TrimSpace(strings.Replace(lines[0], "#", "", -1))

	written, err := time.Parse("*Jan 2, 2006*", lines[1])
	if err != nil {
		return Post{}, fmt.Errorf("Failed to parse post date: %s", err)
	}

	content := strings.Join(lines[2:], "\n\n")

	return Post{title, content, written, false}, nil
}

// NewFSBlogRepository creates a FSBlogRepository in the provided directory. If
// the directory doesn't exist it creates it.
func NewFSBlogRepository(dir string) BlogRepository {
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0700)
	}
	return FSBlogRepository{dir}
}
