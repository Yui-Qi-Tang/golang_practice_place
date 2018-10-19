package http2server

import (
	"net/http"
	"log"
	"fmt"
)

const indexHTML = `<html>
<head>
	<title>Hello World</title>
	<script src="/static/app.js"></script>
	<link rel="stylesheet" href="/static/style.css"">
</head>
<body>
Hello, gopher!
</body>
</html>
`


func CreateServer(config *ServerConfig) {
	log.Printf("start server at %s\n", config.Port)
	// Set static file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(config.Static))))
	// functions
	StartService()
	log.Fatal(
		http.ListenAndServeTLS(config.Port, config.Crt, config.Key, nil),
	)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	pusher, ok := w.(http.Pusher)
	if ok {
		// Push is supported. Try pushing rather than
		// waiting for the browser request these static assets.
		if err := pusher.Push("/static/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
		if err := pusher.Push("/static/style.css", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	fmt.Fprintf(w, indexHTML)
}

func StartService() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
}