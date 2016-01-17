# LitePub

A lightweight static blog generator written in Go.

> Why another one? I wrote a blog post that briefly describes
[why I created it](http://www.mirovarga.com/a-lightweight-static-blog-generator-in-go.html).

## Overview

LitePub is a static blog generator that tries to be as easy to use as possible.

It requires no software dependencies, needs no configuration files, uses no
databases. All it needs is one binary, posts written in
[Markdown](https://en.wikipedia.org/wiki/Markdown) and a set of templates to
render the posts to static HTML files.

Posts don't have to include any special metadata (aka front matter) like title
or date in them - the title, date and optional tags are parsed from
the natural flow of the posts.

## Quick Start

To create a sample blog follow these steps:

1. Download a [release](https://github.com/mirovarga/litepub/releases) and
unpack it to a directory.

2. `cd` to the directory.

3. Create a sample blog:

  	```
  	$ ./litepub create
  	```

4. Build the blog:

	```
	$ ./litepub build
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
  Generating: websites-using-litepub.html
	```

5. Run the built-in server:

	```
	$ ./litepub serve
	Running on http://localhost:2703
	Ctrl+C to quit
	```

6. Open [http://localhost:2703](http://localhost:2703) in your browser.

## Documentation

### Installation

Download a [release](https://github.com/mirovarga/litepub/releases) and unpack
it to a directory. That's all.

> You can optionally add the directory to the `PATH` so you can run `litepub`
from any directory. All examples assume you have `litepub` in your `PATH`.

### Creating a Blog

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

#### Directory Structure

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

#### The **create** Command Reference

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

### Creating Posts

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

#### Draft Posts

Any post can be marked as draft by simply moving it to the `draft` subdirectory
of the `posts` directory. To unmark it just move it back to the `posts` directory.

> Deleting a post is analogous to drafting: just remove it from the `posts`
directory.

### Generating HTML Files for a Blog, aka Building a Blog

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
Generating: websites-using-litepub.html
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

#### The **build** Command Reference

```
Usage:
  litepub build  [<dir>] [-q, --quiet]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -q, --quiet        Show only errors     
```

### Serving a Blog

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

#### Serving a Blog on a Different Port

When starting the server you can specify a port on which to listen with the
`--port` option:

```
$ litepub serve --port 3000
Running on http://localhost:3000
Ctrl+C to quit
```

#### Serving a Blog and Watching for Changes

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

#### Rebuilding a Blog Before Serving

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
Generating: websites-using-litepub.html
Running on http://localhost:2703
Ctrl+C to quit
```

#### The **serve** Command Reference

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

### Templates

The `create` command adds sample templates to the `templates` directory. Of
course, you can change them or create your own from scratch.

LitePub uses the Go [html/template](https://golang.org/pkg/html/template/)
package to define the templates.

> Design changes require no knowledge of Go templates. However changes that
affect what data is displayed will require it less or more (depending on
the change).

#### Structure

There are four required files in the `templates` directory:

```bash
templates/      # the templates and accompanying files (html, css, js, png, etc.)
  layout.tmpl
  index.tmpl
  post.tmpl
  tag.tmpl
```

- `layout.tmpl` defines the common layout for the home page (`index.tmpl`), post
   pages (`post.tmpl`) and tag pages (`tag.tmpl`)
- `index.tmpl` is used when generating the home page (`index.html`)
- `post.tmpl` is used when generating post pages
- and `tag.tmpl` is used when generating tag pages

Besides the four files there can be any number of `html`, `css`, `js`, `png`,
etc. files that are used by the `.tmpl` files.

> If you're not familiar with Go templates, some things in the next sections can
be unclear.

#### Data

Templates have access to data they are meant to display. There are two types of
data: `Post`s and `Tag`s.

##### Posts

A `Post` has the following properties:

- `Title` - the post title
- `Content` - the content of the post as Markdown text
- `Written` - the post's date
- `Tags` - an array of tags the post is tagged with (can be empty)
- `Draft` - `true` if the post is a draft

> To get a post's page URL in a template use the `slug` function (described
below) like this: `<a href="/{{.Title | slug}}.html">A Post</a>`.

##### Tags

A `Tag` has the following properties:

- `Name` - the tag name
- `Posts` - an array of `Post`s that are tagged with the tag sorted by `Written`
  in descending order

> To get a tag's page URL in a template use the `slug` function (described
below) like this: `<a href="/tags/{{.Name | slug}}.html">A Tag</a>`.

The `index.tmpl` template has access to an array of `Post`s sorted by `Written`
in descending order. The `post.tmpl` template has access to the `Post` it
displays. The `tag.tmpl` template has access to the `Tag` it displays.

#### Functions

The `index.tmpl`, `post.tmpl` and `tag.tmpl` templates have access to
the following functions:

##### html

Converts a Markdown string to a raw HTML, for example `{{.Content | html}}`.

##### summary

Extracts the first paragraph of a Markdown string that isn't a header (doesn't
start with a `#`), for example `{{.Content | summary | html}}`.

##### even

Returns `true` if an integer is even, for example
`{{if even $i}}<divclass="row">{{end}}`.

##### inc

Increments an integer by one, for example
`{{if or (eq (inc $i) $l) (not (even $i))}}</div>{{end}}`.

##### slug

Slugifies a string, for example `<a href="/{{.Title | slug}}.html">A Post</a>`.

> The available functions represent my needs when converting my handmade blog
to a generated one.

### Getting Help

To see all available commands and their options use the `--help` option:

```
$ litepub --help
LitePub 0.5.1 [github.com/mirovarga/litepub]
Copyright (c) 2016 Miro Varga [mirovarga.com, hello@mirovarga.com, @mirovarga]

Usage:
  litepub create [<dir>] [-s, --skeleton] [-q, --quiet]
  litepub build  [<dir>] [-q, --quiet]
  litepub serve  [<dir>] [-R, --rebuild] [-p, --port <port>] [-w, --watch] [-q, --quiet]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -s, --skeleton     Don't create sample posts and templates
  -R, --rebuild      Rebuild the blog before serving
  -p, --port <port>  The port to listen on [default: 2703]
  -w, --watch        Rebuild the blog when posts or templates change
  -q, --quiet        Show only errors
  -h, --help         Show this screen
  -v, --version      Show version
```
