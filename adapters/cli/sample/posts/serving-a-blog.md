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

**Next**: [Templates](/templates.html)
