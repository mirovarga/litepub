# Creating a Blog

*Nov 16, 2015*

*Docs, Basics*

The following will create a sample blog in the `litepub-blog` directory:

```
$ litepub create
```

You can override the default directory by providing your own:

```
$ litepub create my-blog
```
> If the directory already exists the command will fail.

If you don't need the sample templates and posts use the `--blank` option:

```
$ litepub create my-blog --blank
```

> Because the template files are required they will be still created but with no
content.

## Directory Structure

Each blog is just a directory with the following structure:

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

**Next**: [Creating Posts](/creating-posts.html)
