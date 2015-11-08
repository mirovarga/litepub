package main

import (
	"fmt"
	"net/http"

	"github.com/go-fsnotify/fsnotify"
)

const defaultPort = "2703"

func server(arguments map[string]interface{}) {
	watch := arguments["--watch"].(int)

	if watch == 1 {
		watcher, _ := fsnotify.NewWatcher()
		defer watcher.Close()

		go func() {
			for {
				select {
				case <-watcher.Events:
					build(nil)
				}
			}
		}()

		watcher.Add(postsDir)
		watcher.Add(templatesDir)
	}

	port, ok := arguments["--port"].([]string)
	if !ok {
		port[0] = defaultPort
	}

	fmt.Printf("Running on http://localhost:%s\n", port[0])
	if watch == 1 {
		fmt.Println("Auto rebuilding when posts or templates change")
	}
	fmt.Println("Ctrl+C to quit")

	http.ListenAndServe(":"+port[0], http.FileServer(http.Dir(outputDir)))
}
