# Serving a Blog

*Nov 13, 2015*

*Docs, Basics*

LitePub has a built-in server so you can see how a generated blog looks like
in a browser. `cd` to the blog's directory and start the server:

```
$ litepub serve
Running on http://localhost:2703
Ctrl+C to quit
```

Now open [http://localhost:2703](http://localhost:2703) in your browser to see
the generated blog.

> Note that the server ignores the draft posts.

## Serving a Blog on a Different Port

When starting the server you can specify a port on which to listen with the
`--port` option:

```
$ litepub serve --port 3000
Running on http://localhost:3000
Ctrl+C to quit
```

## Serving a Blog and Watching for Changes

When creating templates or even writing posts it's quite useful to be able to
immediately see the changes after refreshing the page. To tell LitePub that it
should watch for changes to posts and templates use the `--watch` option:

```
$ litepub serve --watch
Running on http://localhost:2703
Rebuilding when posts or templates change
Ctrl+C to quit
```

> Note that subdirectories in the `posts` and `templates` directories aren't
watched.

## Rebuilding a Blog Before Serving

Sometimes it can be useful to rebuild a blog before serving it, for example when
you don't remember if you made any changes to posts or templates. To rebuild
a blog before serving use the `--rebuild` option:

```
$ litepub serve --rebuild
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
Running on http://localhost:2703
Ctrl+C to quit
```

## The **serve** Command Reference

```
Usage:
  litepub serve  [<dir>] [-R, --rebuild] [-p, --port <port>] [-w, --watch] [-q, --quiet]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -R, --rebuild      Rebuild the blog before serving
  -p, --port <port>  The port to listen on [default: 2703]
  -w, --watch        Rebuild the blog when posts or templates change
  -q, --quiet        Show only errors
```

**Next**: [Templates](/templates.html)
