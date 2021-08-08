package cli

import (
	"net/http"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

const defaultPort = "2703"

func serve(arguments map[string]interface{}) int {
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

	log.Infof("Running on http://localhost:%s\n", port[0])
	if watch == 1 {
		log.Infof("Rebuilding when posts or templates change\n")
	}
	log.Infof("Ctrl+C to quit\n")

	http.ListenAndServe(":"+port[0], http.FileServer(http.Dir(filepath.Join(dir, outputDir))))
	return 0
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
