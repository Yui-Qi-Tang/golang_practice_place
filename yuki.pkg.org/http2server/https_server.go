package http2server

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/websocket"
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

var clients = make(map[*user]bool)

var flag = make(chan bool)

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
	//http.ListenAndServe(config.Port, nil)

}

func StartService() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	http.HandleFunc("/socket/handler", webSocketHandle)
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
	if r.Header.Get("Origin") != "https://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
	}

	// allocate user id
    new_user_id := generateUserId()
	for k, _ := range clients {
		if k.id == new_user_id {
			new_user_id = generateUserId()
			continue
		}
	}
	
	sendFirstJoinMsg(conn, new_user_id)	
	// add new user
	new_user := &user{wsconn: conn, id: new_user_id}
	clients[new_user] = true
	// go echo(conn)
	go chatHandle(new_user)	
}

func sendFirstJoinMsg(conn *websocket.Conn, guess_id int) {
	welcome := &msg{Text: "Hello!!Wellcome join us!!", MyId: guess_id, To: nil, From: nil}
    conn.WriteJSON(welcome)
}

func chatHandle(chater *user) {
	for {
		m := msg{} // custom msg
		err := chater.wsconn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
			chater.wsconn.Close()
			delete(clients, chater)
			fmt.Println(clients)
			flag <- false
		}

		fmt.Printf("Got message: %#v\n", m)
		// board cast msg
		for k, _ := range clients {
			if k.id != chater.id {
				m.From = chater.id
				m.MyId = nil
				
			    err := k.wsconn.WriteJSON(m)
			    if err != nil {
			    	fmt.Printf("send failed!")
			    }
		    }
		}

	}
}

/*
func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
			conn.Close()
			delete(clients, conn)
			fmt.Println(clients)
			flag <- false
		}

		fmt.Printf("Got message: %#v\n", m)

		for k, _ := range clients {
			resp := []byte("You talkin' to me?")
			err := k.WriteMessage(1, resp)
			if err != nil {
				fmt.Printf("send failed!")
			}
		}

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}*/