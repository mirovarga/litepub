package litepub

import "time"

// Blog represents a blog.
type Blog struct {
	ID    string
	Posts []Post
}

// NonDraftPosts returns Posts that are not drafts (they have Draft set to false).
func (b Blog) NonDraftPosts() []Post {
	var posts []Post
	for _, post := range b.Posts {
		if post.Draft {
			continue
		}
		posts = append(posts, post)
	}
	return posts
}

// DraftPosts returns Posts that are drafts (they have Draft set to true).
func (b Blog) DraftPosts() []Post {
	var posts []Post
	for _, post := range b.Posts {
		if !post.Draft {
			continue
		}
		posts = append(posts, post)
	}
	return posts
}

// Post is a Blog's post.
// TODO tags
type Post struct {
	Title string

	// Content of the post (can use Markdown).
	Content string
	Written time.Time
	Draft   bool
}

// BlogRepository stores Blogs.
type BlogRepository interface {
	Store(Blog) error
	FindByID(string) (Blog, error)
}
