package litepub

import "time"

// Blog is a blog.
type Blog struct {
	ID    string
	Posts []Post
}

// Post is a Blog's post.
type Post struct {
	Title string

	// Content of the post (can use Markdown).
	Content string

	// TODO tags

	Written time.Time
}

// BlogRepository stores Blogs.
type BlogRepository interface {
	Store(Blog) error
	FindByID(string) (Blog, error)
}
