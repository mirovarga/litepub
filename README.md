# LitePub

A lightweight static blog generator written in Go.

> Why another one? I wrote a
[blog post](http://www.mirovarga.com/a-lightweight-static-blog-generator-in-go.html)
that briefly describes it.

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
www/            # the generated HTML files (plus copied accompanying files)
```

### Creating Posts

To create a post just add a [Markdown](http://daringfireball.net/projects/markdown/)
file in the `posts` directory. The file name and extension aren't important,
only the content of the file.

> All posts need to be stored directly in the `posts` directory. In other words,
subdirectories in the `posts` directory are ignored when looking for posts.

Each post looks like this (it's an
[actual post](http://www.mirovarga.com/building-an-event-store-in-node-js.html)
from my blog):

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

- Line `1` is the post's title. If it starts with one or more `#`s they are
stripped. So in this case the title becomes *Building an Event Store in Node.js*.
- Line `3` is the post's date. It has to be in the `*MMM d, YYYY*` format.
- The rest is the content of the post.

> Both the post's title and date are required.

#### Draft Posts

Any post can be marked as draft by simply moving it to the `draft` subdirectory
of the `posts` directory. To unmark it just move it back to the `posts` directory.

> Deleting a post is analogous to drafting: just remove it from the `posts`
directory.

### Generating HTML files for a Blog, aka Building a Blog

To generate the HTML files for a blog `cd` to the blog's directory and use the
`build` command:

```
$ litepub build
Generating: index.html
Generating: welcome-to-litepub.html
```

> The draft posts are ignored when building a blog.

LitePub takes the `*.tmpl` files from the `templates` directory, applies them to
posts stored in the `posts` directory and generates the HTML files to the `www`
directory. It also copies all accompanying files (and directories) from
the `templates` directory to the `www` directory.

> The generated HTML files' names are created by slugifying the post title and
adding the `html` extension. For example, a post with the
*Building an Event Store in Node.js* title is generated to the
`building-an-event-store-in-node-js.html` file.

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

There are three required files in the `templates` directory:

```sh
templates/      # the templates and accompanying files (html, css, js, png, etc.)
  layout.tmpl
  index.tmpl
  post.tmpl
```

- `layout.tmpl` defines the common layout for both the home page (`index.tmpl`)
and post pages (`post.tmpl`)
- `index.tmpl` is used to generate the home page (`index.html`)
- and `post.tmpl` is used to generate post pages

Besides the three files there can be any number of `html`, `css`, `js`, `png`,
etc. files that are used by the `.tmpl` files.

#### Objects

There is only one type of object - the `post`:

Each post has the following properties:
- `Title` - as plain text
- `Content` - everything except `Title` and `Written` as Markdown text
- `Written` - the post's date
- `Slug` - generated from the `Title`

The `index.tmpl` template has access to an array of all `post`s sorted by
`Written` in descending order. The `post.tmpl` template has access to the `post`
it represents.

#### Functions

Both the `index.tmpl` and `post.tmpl` templates have access to the following
functions:

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
