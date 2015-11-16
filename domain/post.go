package domain

import (
	"sort"
	"time"
)

// Post is a Blog's post.
type Post struct {
	Title string

	// Content of the post (can use Markdown).
	Content string
	Written time.Time
	Tags    []string
	Draft   bool
}

type sortablePosts []Post

func (p sortablePosts) Len() int {
	return len(p)
}

func (p sortablePosts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p sortablePosts) Less(i, j int) bool {
	return p[i].Written.Before(p[j].Written)
}

func sortByDate(posts []Post, asc bool) {
	if asc {
		sort.Sort(sortablePosts(posts))
	} else {
		sort.Sort(sort.Reverse(sortablePosts(posts)))
	}
}

func filterByTags(posts []Post, tags ...string) []Post {
	if len(tags) == 0 {
		return posts
	}

	filteredPosts := []Post{}
	for _, post := range posts {
		for _, postTag := range post.Tags {
			for _, tag := range tags {
				if postTag == tag {
					filteredPosts = append(filteredPosts, post)
				}
			}
		}
	}
	return filteredPosts
}
