package main

// TODO -q, --quiet for each command (own log that can be set to quiet mode)
// TODO docs for <dir>
const usage = `
LitePub - a lightweight static blog generator, http://litepub.com

Usage:
  litepub create [<dir>] [-b, --blank]
  litepub build  [<dir>]
  litepub serve  [<dir>] [-p, --port <port>] [-w, --watch]

Arguments:
  <dir>  The directory to create the blog in or look for; it will be created if
         it doesn't exist (only when creating a blog) [default: .]

Options:
  -b, --blank        Don't create sample posts and templates
  -p, --port <port>  The port to listen on [default: 2703]
  -w, --watch        Rebuild the blog when posts or templates change
  -h, --help         Show this screen
  -v, --version      Show version
`
