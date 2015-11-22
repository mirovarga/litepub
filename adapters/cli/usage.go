package main

// TODO -d, --dir <dir> option for each command so it can be run from any directory
// TODO -q, --quiet for each command
const usage = `
LitePub - a lightweight static blog generator, http://litepub.com

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
`
