package litepub

import (
	"fmt"
	"time"
)

// Authors wrap the use cases that authors of Blogs can apply on the domain.
type Authors struct {
	blogRepository BlogRepository
}

// CreateBlog creates a Blog with the ID.
//
// If another Blog with the same ID already exists it returns an error.
func (a Authors) CreateBlog(id string) error {
	_, err := a.blogRepository.FindByID(id)
	if err == nil {
		return fmt.Errorf("Blog already exists: %s", id)
	}

	return a.blogRepository.Store(Blog{ID: id})
}

// CreatePost creates a Post and adds it to the Blog with the ID.
//
// If the Blog with the ID doesn't exist it returns an error.
func (a Authors) CreatePost(id, title, content string, written time.Time) error {
	return a.createPost(id, title, content, written, false)
}

// CreateDraftPost creates a Post, sets its Draft to true and adds it to the
// Blog with the ID.
//
// If the Blog with the ID doesn't exist it returns an error.
func (a Authors) CreateDraftPost(id, title, content string, written time.Time) error {
	return a.createPost(id, title, content, written, true)
}

func (a Authors) createPost(id, title, content string, written time.Time,
	draft bool) error {
	blog, err := a.blogRepository.FindByID(id)
	if err != nil {
		return err
	}

	blog.Posts = append(blog.Posts, Post{title, content, written, draft})
	return a.blogRepository.Store(blog)
}

// NewAuthors creates Authors that will use the provided repository for
// authoring Blogs.
func NewAuthors(blogRepository BlogRepository) Authors {
	return Authors{blogRepository}
}

// Readers wrap the use cases that readers of Blogs can apply on the domain.
type Readers struct {
	blogRepository BlogRepository
}

// GetBlog gets the Blog with the ID.
//
// If the Blog with the ID doesn't exist it returns an error.
func (r Readers) GetBlog(id string) (Blog, error) {
	return r.blogRepository.FindByID(id)
}

// NewReaders creates Readers that will use the provided repository for
// accessing Blogs.
func NewReaders(blogRepository BlogRepository) Readers {
	return Readers{blogRepository}
}
