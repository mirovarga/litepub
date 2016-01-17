# Generating HTML Files for a Blog, aka Building a Blog

*Nov 14, 2015*

*Docs, Basics*

To generate the HTML files for a blog `cd` to the blog's directory and use the
`build` command:

```
$ litepub build
Generating: index.html
Generating: tags/reference.html
Generating: tags/tutorial.html
Generating: tags/advanced.html
Generating: tags/docs.html
Generating: tags/basics.html
Generating: overview.html
Generating: quick-start.html
Generating: installation.html
Generating: creating-a-blog.html
Generating: creating-posts.html
Generating: generating-html-files-for-a-blog-aka-building-a-blog.html
Generating: serving-a-blog.html
Generating: templates.html
Generating: getting-help.html
```

> The draft posts are ignored when building a blog.

LitePub takes the `*.tmpl` files from the `templates` directory, applies them to
posts stored in the `posts` directory and generates the HTML files to the `www`
directory. It also copies all accompanying files (and directories) from
the `templates` directory to the `www` directory.

> The generated HTML file names are created by slugifying the post title (or
the tag name when generating tag pages) and adding the `html` extension. For
example, a post with the *How I Switched from Java to JavaScript* title is
generated to the `how-i-switched-from-java-to-javascript.html` file.

## The **build** Command Reference

```
Usage:
  litepub build  [<dir>] [-q, --quiet]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -q, --quiet        Show only errors     
```

**Next**: [Serving a Blog](/serving-a-blog.html)
