# Getting Help

*Nov 11, 2015*

*Docs, Reference*

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
