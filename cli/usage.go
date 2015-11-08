package main

const usage = `
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
`
