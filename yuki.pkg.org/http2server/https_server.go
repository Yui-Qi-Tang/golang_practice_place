package http2server

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/websocket"
	// "html/template"
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

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type msg struct {
	Num int
}


func CreateServer(config *ServerConfig) {
	log.Printf("start server at %s\n", config.Port)
	// Set static file
	// filter   "prefix" of the request and leave
	//            request                          "prefix"                      file here
	http.Handle("/chatroom/", http.StripPrefix("/chatroom/", http.FileServer(http.Dir(config.Static))))
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
	fmt.Println(r.UserAgent())

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

func webSocketHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
	}
	go echo(conn)
}

func StartService() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	http.HandleFunc("/socket/handler", webSocketHandle)
}

func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}