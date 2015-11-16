package application

import (
	"fmt"
	"time"

	"mirovarga.com/litepub/domain"
)

// Authors wrap the use cases for authors of Blogs.
type Authors struct {
	blogRepository domain.BlogRepository
}

// CreateBlog creates a Blog with the ID.
//
// If another Blog with the same ID already exists it returns an error.
func (a Authors) CreateBlog(id string) error {
	_, err := a.blogRepository.FindByID(id)
	if err == nil {
		return fmt.Errorf("Blog already exists: %s", id)
	}

	return a.blogRepository.Store(domain.Blog{ID: id})
}

// CreatePost creates a Post and adds it to the Blog with the ID.
//
// If the Blog with the ID doesn't exist it returns an error.
func (a Authors) CreatePost(id, title, content string, written time.Time,
	tags ...string) error {
	return a.createPost(id, title, content, written, tags, false)
}

// CreateDraftPost creates a Post, sets its Draft to true and adds it to the
// Blog with the ID.
//
// If the Blog with the ID doesn't exist it returns an error.
func (a Authors) CreateDraftPost(id, title, content string, written time.Time,
	tags ...string) error {
	return a.createPost(id, title, content, written, tags, true)
}

func (a Authors) createPost(id, title, content string, written time.Time,
	tags []string, draft bool) error {
	blog, err := a.blogRepository.FindByID(id)
	if err != nil {
		return err
	}

	blog.Posts = append(blog.Posts, domain.Post{
		Title:   title,
		Content: content,
		Written: written,
		Tags:    tags,
		Draft:   draft,
	})
	return a.blogRepository.Store(blog)
}

// NewAuthors creates Authors that will use the provided repository for
// authoring Blogs.
func NewAuthors(blogRepository domain.BlogRepository) Authors {
	return Authors{blogRepository}
}

// Readers wrap the use cases for readers of Blogs.
type Readers struct {
	blogRepository domain.BlogRepository
}

// GetBlog gets the Blog with the ID.
//
// If the Blog with the ID doesn't exist it returns an error.
func (r Readers) GetBlog(id string) (domain.Blog, error) {
	return r.blogRepository.FindByID(id)
}

// NewReaders creates Readers that will use the provided repository for
// accessing Blogs.
func NewReaders(blogRepository domain.BlogRepository) Readers {
	return Readers{blogRepository}
}
