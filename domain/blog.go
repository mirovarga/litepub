package domain

// Blog represents a blog.
type Blog struct {
	ID    string
	Posts []Post
}

// PostsByDate returns Posts of the Blog sorted in ascending or descending order
// (if asc == false). If includeDrafts == true draft Posts are also included.
// If withTags is present only Posts having the tags are included.
func (b Blog) PostsByDate(asc, includeDrafts bool, withTags ...string) []Post {
	posts := []Post{}
	for _, post := range b.Posts {
		if post.Draft && !includeDrafts {
			continue
		}
		posts = append(posts, post)
	}

	sortByDate(posts, asc)
	return filterByTags(posts, withTags...)
}

// Tags returns all tags used in Posts of the Blog. If lookInDrafts == true
// draft Posts are also checked.
func (b Blog) Tags(lookInDrafts bool) []string {
	tags := make(map[string]string)
	for _, post := range b.Posts {
		if post.Draft && !lookInDrafts {
			continue
		}
		for _, tag := range post.Tags {
			tags[tag] = tag
		}
	}

	blogTags := make([]string, len(tags))
	i := 0
	for tag := range tags {
		blogTags[i] = tag
		i++
	}
	return blogTags
}

// BlogRepository persists Blogs.
type BlogRepository interface {
	// Store adds or updates a Blog in the repository.
	Store(Blog) error

	FindByID(string) (Blog, error)
}
