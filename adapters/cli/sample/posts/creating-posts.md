# Creating Posts

*Nov 15, 2015*

*Docs, Basics*

To create a post just add a [Markdown](https://en.wikipedia.org/wiki/Markdown)
file in the `posts` directory. The file name and extension aren't important,
only the content of the file.

> All posts need to be stored directly in the `posts` directory. In other words,
subdirectories in the `posts` directory are ignored when looking for posts.

Each post looks like this (it's the start of an
[actual post](http://www.mirovarga.com/how-i-switched-from-java-to-javascript.html)
from my blog):

```markdown
1 # How I Switched from Java to JavaScript
2
3 *Jan 25, 2015*
4
5 *Java, JavaScript*
6
7 I know that there are lots of posts about why JavaScript, or more specifically
8 Node.js, is better than Java but nevertheless I wanted to contribute, too.
9 ...
```

- Line `1` is the post's title. If it starts with one or more `#`s they are
stripped. So in this case the title becomes *How I Switched from Java to JavaScript*.
- Line `3` is the post's date. It has to be in the `*MMM d, YYYY*` format.
- Line `5` are comma separated post tags.
- Anything below line `6` is the content of the post.

> The post's title and date are required. Tags are optional.

## Draft Posts

Any post can be marked as draft by simply moving it to the `draft` subdirectory
of the `posts` directory. To unmark it just move it back to the `posts` directory.

> Deleting a post is analogous to drafting: just remove it from the `posts`
directory.

**Next**: [Generating HTML Files for a Blog, aka Building a Blog](/generating-html-files-for-a-blog-aka-building-a-blog.html)
