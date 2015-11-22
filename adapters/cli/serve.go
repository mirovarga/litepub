package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"gopkg.in/fsnotify.v1"
)

const defaultPort = "2703"

// TODO -o, --output <dir>  Generate the blog to the specified directory [default: www]
func serve(arguments map[string]interface{}) {
	dir := arguments["<dir>"].(string)

	if arguments["--rebuild"].(int) == 1 {
		build(map[string]interface{}{"<dir>": dir})
	}

	port, ok := arguments["--port"].([]string)
	if !ok {
		port[0] = defaultPort
	}

	watch := arguments["--watch"].(int)

	if watch == 1 {
		go watchDirs(dir)
	}

	fmt.Printf("Running on http://localhost:%s\n", port[0])
	if watch == 1 {
		fmt.Println("Rebuilding when posts or templates change")
	}
	fmt.Println("Ctrl+C to quit")

	http.ListenAndServe(":"+port[0], http.FileServer(http.Dir(filepath.Join(dir, outputDir))))
}

func watchDirs(dir string) {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	watcher.Add(filepath.Join(dir, postsDir))
	watcher.Add(filepath.Join(dir, templatesDir))

	for {
		select {
		case <-watcher.Events:
			build(map[string]interface{}{"<dir>": dir})
		}
	}
}
