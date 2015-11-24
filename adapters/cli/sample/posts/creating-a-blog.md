# Creating a Blog

*Nov 16, 2015*

*Docs, Basics*

The following will create a sample blog in the current directory:

```
$ litepub create
```

If you don't need the sample templates and posts use the `--skeleton` option:

```
$ litepub create --skeleton
```

> Because the template files are required they will be still created but with no
content.

## Directory Structure

Each blog is stored in a directory with the following structure:

```bash
posts/          # the posts
  draft/        # the draft posts (they are ignored when building the blog)
templates/      # the templates and accompanying files (html, css, js, png, etc.)
  layout.tmpl
  index.tmpl
  post.tmpl
  tag.tmpl
www/            # the generated HTML files (plus copied accompanying files)
```

## The **create** Command Reference

```
Usage:
  litepub create [<dir>] [-s, --skeleton] [-q, --quiet]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -s, --skeleton     Don't create sample posts and templates
  -q, --quiet        Show only errors
```

**Next**: [Creating Posts](/creating-posts.html)
