/*
   TODO: better to set the static files(server key/crt files, css files, js files and html files)
*/
package serverpush

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

func Demonstration() {
	flag.Parse()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("http2server/static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})
	log.Fatal(http.ListenAndServeTLS(*httpAddr, "http2server/server.crt", "http2server/server.key", nil))
}

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
