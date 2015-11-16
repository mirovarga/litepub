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

Posts don't have to include any special metadata like title or date in them;
of course they do have titles, dates and optionally tags but they flow naturally
and posts thus look like posts on their own.

```markdown
# How I Switched from Java to JavaScript

*Jan 25, 2015*

*Java, JavaScript*

I know that there are lots of posts about why JavaScript, or more specifically
Node.js, is better than Java but nevertheless I wanted to contribute, too.
```

LitePub supports tagging posts, draft posts and provides a built-in HTTP server
that can rebuild a blog on the fly when its posts or templates change.

## Quick Start

1. Download a [release](https://github.com/mirovarga/litepub/releases) and
unpack it to a directory.

2. `cd` to the directory.

3. Create a sample blog:

  ```
  $ ./litepub create
  ```

4. Build the blog:

  ```
  $ cd litepub-blog    
  $ ../litepub build
  Generating: index.html
  Generating: tags/litepub.html
  Generating: welcome-to-litepub.html
  ```

5. Run the built-in server:

  ```
  $ ../litepub server
  Running on http://localhost:2703
  Ctrl+C to quit
  ```

6. Open [http://localhost:2703](http://localhost:2703) in your browser.

## Documentation

### Installation

Download a [release](https://github.com/mirovarga/litepub/releases) and unpack
it to a directory. That's all.

> You can optionally add the directory to the `PATH` so you can run `litepub`
from any directory. All examples below assume you have `litepub` in your `PATH`.

### Creating a Blog

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

#### Directory Structure

Each blog is just a directory with the following structure:

```sh
posts/          # the posts
  draft/        # the draft posts (they are ignored when building the blog)
templates/      # the templates and accompanying files (html, css, js, png, etc.)
  layout.tmpl
  index.tmpl
  post.tmpl
  tag.tmpl
www/            # the generated HTML files (plus copied accompanying files)
```

### Creating Posts

To create a post just add a [Markdown](https://en.wikipedia.org/wiki/Markdown)
file in the `posts` directory. The file name and extension aren't important,
only the content of the file.

> All posts need to be stored directly in the `posts` directory. In other words,
subdirectories in the `posts` directory are ignored when looking for posts.

Each post looks like this (it's an
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
Generating: tags/litepub.html
Generating: welcome-to-litepub.html
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

### Serving a Blog

LitePub has a built-in server so you can see how a generated blog looks like
in a browser. `cd` to the blog's directory and start the server:

```
$ litepub server
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
$ litepub server --port 3000
Running on http://localhost:3000
Ctrl+C to quit
```

#### Serving a Blog and Watching for Changes

When creating templates or even writing posts it's quite useful to be able to
immediately see the changes after refreshing the page. To tell LitePub that it
should watch for changes to posts and templates use the `--watch` option:

```
$ litepub server --watch
Running on http://localhost:2703
Rebuilding when posts or templates change
Ctrl+C to quit
```

> Note that subdirectories in the `posts` and `templates` directories aren't
watched.

### Templates

The `create` command adds sample templates to the `templates` directory. However,
they are very basic and serve only the purpose to get you started quickly.

Of course, you can change the sample templates or create your own from scratch.
LitePub uses the Go [html/template](https://golang.org/pkg/html/template/)
package to define the templates.

> If you're not familiar with the Go `html/template` package, some things in
the next sections can be unclear.

There are four required files in the `templates` directory:

```sh
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

#### Data

Templates have access to data they are meant to display. There are two types of
data: `post`s and `tag`s.

A `post` has the following properties:

- `Title` - the post title
- `Content` - the content of the post as Markdown text
- `Written` - the post's date
- `Tags` - an array of tags the post is tagged with (can be empty)
- `Draft` - `true` if the post is a draft

> To get a post's page URL in a template use the `slug` function (described
below) like this: `<a href="/{{.Title | slug}}.html">A Post</a>`.

A `tag` has the following properties:

- `Name` - the tag name
- `Posts` - an array of `post`s that are tagged with the tag sorted by `Written`
  in descending order

> To get a tag's page URL in a template use the `slug` function (described
below) like this: `<a href="/tags/{{.Name | slug}}.html">A Tag</a>`.

The `index.tmpl` template has access to an array of `post`s sorted by `Written`
in descending order. The `post.tmpl` template has access to the `post` it
displays. The `tag.tmpl` template has access to the `tag` it displays.

#### Functions

The `index.tmpl`, `post.tmpl` and `tag.tmpl` templates have access to
the following functions:

##### `html`

Converts a Markdown string to a raw HTML, for example `{{.Content | html}}`.

##### `summary`

Extracts the first paragraph of a Markdown string that isn't a header (doesn't
start with a `#`), for example `{{.Content | summary | html}}`.

##### `even`

Returns `true` if an integer is even, for example
`{{if even $i}}<divclass="row">{{end}}`.

##### `inc`

Increments an integer by one, for example
`{{if or (eq (inc $i) $l) (not (even $i))}}</div>{{end}}`.

##### `slug`

Slugifies a string, for example `<a href="/{{.Title | slug}}.html">A Post</a>`.

> The available functions represent my needs when converting my handmade blog
to a generated one.

### Getting Help

To see all available commands and their options use the `--help` option:

```
$ litepub --help
LitePub - a lightweight static blog generator, https://github.com/mirovarga/litepub

Usage:
  litepub create [<name>] [-b, --blank]
  litepub build
  litepub server [-p, --port <port>] [-w, --watch]

Arguments:
  <name>             Name of the blog [default: litepub-blog]

Options:
  -b, --blank        Don't create sample posts and templates
  -p, --port <port>  The port to listen on [default: 2703]
  -w, --watch        Rebuild the blog when posts or templates change
  -h, --help         Show this screen
  -v, --version      Show version
```

## Websites Using LitePub

The only website using LitePub that I know of is
[my blog](http://www.mirovarga.com).

> If you happen to know of any other websites please
[send me an email](mailto:hello@mirovarga.com?subject=Another%20LitePub%20Website)
and I'll add them here.

## Implementation Notes

The project tries to apply
the [hexagonal](http://alistair.cockburn.us/Hexagonal+architecture),
aka [clean](http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html),
aka [onion](http://jeffreypalermo.com/blog/the-onion-architecture-part-1/)
architecture. For example, that's the reason why the CLI source is under
the `adapters` directory.

There's a nice post about
[applying the architecture in Go](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/).
