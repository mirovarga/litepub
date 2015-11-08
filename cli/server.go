package main

import (
	"fmt"
	"net/http"
)

const defaultPort = "2703"

// TODO -w, --watch  Auto rebuild the blog when posts or templates change
func server(arguments map[string]interface{}) {
	port, ok := arguments["--port"].([]string)
	if !ok {
		port[0] = defaultPort
	}

	fmt.Printf("Running on http://localhost:%s\nCtrl+C to quit\n", port[0])
	http.ListenAndServe(":"+port[0], http.FileServer(http.Dir(outputDir)))
}
