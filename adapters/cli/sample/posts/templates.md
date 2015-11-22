# Templates

*Nov 12, 2015*

*Docs, Advanced*

The `create` command adds sample templates to the `templates` directory. Of
course, you can change them or create your own from scratch.

LitePub uses the Go [html/template](https://golang.org/pkg/html/template/)
package to define the templates.

> Design changes require no knowledge of Go templates. However changes that
affect what data is displayed will require it less or more (depending on
the change).

## Structure

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

## Data

Templates have access to data they are meant to display. There are two types of
data: `Post`s and `Tag`s.

### Posts

A `Post` has the following properties:

- `Title` - the post title
- `Content` - the content of the post as Markdown text
- `Written` - the post's date
- `Tags` - an array of tags the post is tagged with (can be empty)
- `Draft` - `true` if the post is a draft

> To get a post's page URL in a template use the `slug` function (described
below) like this: `<a href="/{{.Title | slug}}.html">A Post</a>`.

### Tags

A `Tag` has the following properties:

- `Name` - the tag name
- `Posts` - an array of `Post`s that are tagged with the tag sorted by `Written`
  in descending order

> To get a tag's page URL in a template use the `slug` function (described
below) like this: `<a href="/tags/{{.Name | slug}}.html">A Tag</a>`.

The `index.tmpl` template has access to an array of `Post`s sorted by `Written`
in descending order. The `post.tmpl` template has access to the `Post` it
displays. The `tag.tmpl` template has access to the `Tag` it displays.

## Functions

The `index.tmpl`, `post.tmpl` and `tag.tmpl` templates have access to
the following functions:

### html

Converts a Markdown string to a raw HTML, for example `{{.Content | html}}`.

### summary

Extracts the first paragraph of a Markdown string that isn't a header (doesn't
start with a `#`), for example `{{.Content | summary | html}}`.

### even

Returns `true` if an integer is even, for example
`{{if even $i}}<divclass="row">{{end}}`.

### inc

Increments an integer by one, for example
`{{if or (eq (inc $i) $l) (not (even $i))}}</div>{{end}}`.

### slug

Slugifies a string, for example `<a href="/{{.Title | slug}}.html">A Post</a>`.

> The available functions represent my needs when converting my handmade blog
to a generated one.

**Next**: [Getting Help](/getting-help.html)
