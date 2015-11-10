# LitePub

A lightweight static blog generator written in Go.

## Quick start

1. Download a [binary](https://github.com/mirovarga/litepub/tree/master/bin).

2. Create a sample blog:

  ```
  $ litepub create
  ```

3. Build the blog:

  ```
  $ cd litepub-blog
  $ ../litepub build
  ```

4. Run the built-in server:

  ```
  $ ../litepub server
  ```

5. Open [http://localhost:2703](http://localhost:2703) in your browser

## Usage

```
LitePub - a lightweight static blog generator

Usage:
  litepub create [<name>] [-b, --blank]
  litepub build
  litepub server [-p, --port <port>] [-w, --watch]

Arguments:
  <name>             Name of the blog [default: litepub-blog]

Options:
  -b, --blank        Don't create sample posts and templates
  -p, --port <port>  The port to listen on [default: 2703]
  -w, --watch        Auto rebuild the blog when posts or templates change
  -h, --help         Show this screen
  -v, --version      Show version

```

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
- `Title` - as plain text
- `Content` - everything except `Title` and `Written` as Markdown text
- `Written` - the post's date
- `Slug` - generated from the `Title`

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
[blog](http://www.mirovarga.com/building-an-event-store-in-node-js.html):

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

## Implementation notes

The project tries to apply
the [hexagonal](http://alistair.cockburn.us/Hexagonal+architecture),
aka [clean](http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html),
aka [onion](http://jeffreypalermo.com/blog/the-onion-architecture-part-1/)
architecture. For example, that's the reason why the CLI source is under
the `adapters` directory.

There's a thorough post about
[applying the architecture in Go](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/).
