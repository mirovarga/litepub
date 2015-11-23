# Quick Start

*Nov 18, 2015*

*Docs, Tutorial*

To create a sample blog follow these steps:

1. Download a [release](https://github.com/mirovarga/litepub/releases) and
unpack it to a directory.

2. `cd` to the directory.

3. Create a sample blog:

  	```
  	$ ./litepub create
  	```

4. Build the blog:

	```
	$ ./litepub build
  Generating: index.html
  Generating: tags/reference.html
  Generating: tags/tutorial.html
  Generating: tags/advanced.html
  Generating: tags/docs.html
  Generating: tags/basics.html
  Generating: overview.html
  Generating: quick-start.html
  Generating: installation.html
  Generating: creating-a-blog.html
  Generating: creating-posts.html
  Generating: generating-html-files-for-a-blog-aka-building-a-blog.html
  Generating: serving-a-blog.html
  Generating: templates.html
  Generating: getting-help.html
  Generating: websites-using-litepub.html
	```

5. Run the built-in server:

	```
	$ ./litepub serve
	Running on http://localhost:2703
	Ctrl+C to quit
	```

6. Open [http://localhost:2703](http://localhost:2703) in your browser.

**Next**: [Installation](/installation.html)
