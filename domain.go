package litepub

import "time"

type Blog struct {
	ID    string
	Posts []Post
}

// XXX what about this?
// type Markdown string
type Post struct {
	Title   string
	Content string
	// TODO tags
	Written time.Time
}

type BlogRepository interface {
	Store(Blog) error
	FindByID(string) (Blog, error)
}
