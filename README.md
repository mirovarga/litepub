# LitePub

A lightweight static blog generator written in Go.

## Quick start

Download a [binary](https://github.com/mirovarga/litepub/tree/master/bin):

> Currently there's only a Linux x64 binary (tested on Ubuntu 14.04) because of
an OS specific code (will be fixed soon).

```
$ wget https://github.com/mirovarga/litepub/raw/master/bin/linux/x64/litepub
$ chmod +x litepub
```

Add `litepub` to the `PATH`:

```
$ export PATH=$PATH:/path/to/where/you/downloaded/the/release
```

Clone the [source of my blog](https://github.com/mirovarga/mirovarga.com) so you
have something to start with:

```
$ git clone https://github.com/mirovarga/mirovarga.com
```

`cd` to the directory where you cloned the blog and run:

```
$ litepub generate
```

The blog wil be generated to the `www` directory.

## Documentation

### How it works

LitePub takes the `*.tmpl` files from the `templates` directory, applies them to
posts stored in the `posts` directory and generates the blog to the `www`
directory.

### Templates

Templates are stored in the `templates` directory.

LitePub uses the Go [html/template](https://golang.org/pkg/html/template/)
package to define templates for blogs.

There are three required files in the `templates` directory:
- `layout.tmpl`
- `index.tmpl`
- `post.tmpl`

`layout.tmpl` defines the common layout for both the home page (`index.tmpl`)
and the post page (`post.tmpl`).

Besides the three files there can be any number of `css`, `js`, `png`, ... files
that are used by the `.tmpl` files.

#### Objects

There is only one type of object - the `post`:

Each post has the following properties:
- `Title` (plain text)
- `Written` (the post's date)
- `Content` (Markdown)
- `Slug` (generated from the `Title`)

```go
type post struct {
	// Plain text
	Title string

	// The post's date
	Written time.Time

	// Markdown
	Content string

	// Generated from the title
	Slug string
}
```

The `index.tmpl` template has access to an array of all `post`s sorted by
`Written` in descending order.

The `post.tmpl` template has access to the `post` it represents.

#### Functions

> The available functions represent my needs when converting my handmade blog
to a generated one.

Both the `index.tmpl` and `post.tmpl` templates have access to the following
functions:

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

### Posts

Posts are stored in the `posts` directory.

#### Format

LitePub posts are written in Markdown.

The first line of a post is its title. If it starts with one or more `#`s they
are ignored.

The third line is the post's date. It is in the `*MMM d, YYYY*` format.

The rest is the content of the post.

> Both title and date are required.

An example post taken from my
[blog](http://www.mirovarga.com/building-an-event-store-in-node-js):

```markdown
1 # Building an Event Store in Node.js
2
3 *Jan 21, 2015*
4
5 As I quite like the idea of
6 [event sourcing](http://docs.geteventstore.com/introduction/event-sourcing-basics)
7 I decided to build a simple event store in Node.js.
8 ...
```

## Usage

```
LitePub - a lightweight static blog generator

Usage:
  litepub generate

Options:
  -h, --help      Show this screen
  -v, --version   Show version
```

To generate a blog, `cd` to its directory and run:

```
$ litepub generate
```

The blog wil be generated to the `www` directory.
