package lib

import (
	"sort"
	"time"
)

// Blog is just a collection of Posts.
type Blog []Post

// PostsByDate returns Posts of the Blog sorted in ascending or descending order
// (if asc == false). If includeDrafts == true draft Posts are also included.
// If withTags is present only Posts having the tags are included.
func (b Blog) PostsByDate(asc, includeDrafts bool, withTags ...string) []Post {
	sortByDate(b, asc)

	var posts []Post
	for _, post := range b {
		if post.Draft && !includeDrafts {
			continue
		}
		posts = append(posts, post)
	}

	return filterByTags(posts, withTags...)
}

// Tags returns all tags used in Posts of the Blog. If lookInDrafts == true
// draft Posts are also checked.
func (b Blog) Tags(lookInDrafts bool) []string {
	tags := make(map[string]string)
	for _, post := range b {
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

func (b Blog) Len() int {
	return len(b)
}

func (b Blog) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Blog) Less(i, j int) bool {
	return b[i].Written.Before(b[j].Written)
}

// Post is a Blog's post.
type Post struct {
	Title string

	// Content of the post (can use Markdown).
	Content string
	Written time.Time
	Tags    []string
	Draft   bool
}

func sortByDate(blog Blog, asc bool) {
	if asc {
		sort.Sort(blog)
	} else {
		sort.Sort(sort.Reverse(blog))
	}
}

func filterByTags(blog Blog, tags ...string) []Post {
	if len(tags) == 0 {
		return blog
	}

	var posts []Post
	for _, post := range blog {
		for _, postTag := range post.Tags {
			for _, tag := range tags {
				if postTag == tag {
					posts = append(posts, post)
				}
			}
		}
	}
	return posts
}
