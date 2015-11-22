package main

import (
	"fmt"
	"net/http"

	"gopkg.in/fsnotify.v1"
)

const defaultPort = "2703"

// TODO -o, --output <dir>  Generate the blog to the specified directory [default: www]
// TODO -R, --rebuild  Rebuild the blog before serving
func serve(arguments map[string]interface{}) {
	port, ok := arguments["--port"].([]string)
	if !ok {
		port[0] = defaultPort
	}

	watch := arguments["--watch"].(int)

	if watch == 1 {
		go watchDirs()
	}

	fmt.Printf("Running on http://localhost:%s\n", port[0])
	if watch == 1 {
		fmt.Println("Rebuilding when posts or templates change")
	}
	fmt.Println("Ctrl+C to quit")

	http.ListenAndServe(":"+port[0], http.FileServer(http.Dir(outputDir)))
}

func watchDirs() {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	watcher.Add(postsDir)
	watcher.Add(templatesDir)

	for {
		select {
		case <-watcher.Events:
			build(nil)
		}
	}
}
