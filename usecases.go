package litepub

import (
	"fmt"
	"time"
)

func NewUsecases(blogRepository BlogRepository) Usecases {
	return Usecases{blogRepository}
}

type Usecases struct {
	blogRepository BlogRepository
}

func (u Usecases) CreateBlog(id string) error {
	_, err := u.GetBlog(id)
	if err == nil {
		return fmt.Errorf("Blog already exists: %s", id)
	}

	return u.blogRepository.Store(Blog{ID: id})
}

func (u Usecases) CreatePost(id, title, content string, written time.Time) error {
	blog, err := u.GetBlog(id)
	if err != nil {
		return err
	}

	blog.Posts = append(blog.Posts, Post{title, content, written})
	return u.blogRepository.Store(blog)
}

func (u Usecases) GetBlog(id string) (Blog, error) {
	return u.blogRepository.FindByID(id)
}
