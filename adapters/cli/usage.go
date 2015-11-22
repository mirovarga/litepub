package main

// TODO -q, --quiet for each command (own log that can be set to quiet mode)
// TODO docs for <dir>, --rebuild, --skeleton
const usage = `
LitePub - a lightweight static blog generator, http://litepub.com

Usage:
  litepub create [<dir>] [-s, --skeleton]
  litepub build  [<dir>]
  litepub serve  [<dir>] [-R, --rebuild] [-p, --port <port>] [-w, --watch]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -s, --skeleton     Don't create sample posts and templates
  -R, --rebuild      Rebuild the blog before serving
  -p, --port <port>  The port to listen on [default: 2703]
  -w, --watch        Rebuild the blog when posts or templates change
  -h, --help         Show this screen
  -v, --version      Show version
`
