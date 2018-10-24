package main

import 	"yuki.pkg.org/httpserver"
import "fmt"

func test(w httpserver.Resp, r httpserver.Req) {
	fmt.Fprintf(w, "Hello world!")
}

func test2(w httpserver.Resp, r httpserver.Req) {
	fmt.Fprintf(w, "Hello world2!")
}


func main() {
	// Pass test function to SetHandleFunc
	httpserver.SetHandleFunc("/test", test)
	httpserver.SetHandleFunc("/test2", test2)
	httpConfig := &httpserver.HTTPConfig{Port: ":8001", StaticFilePath: "./test"}
	simpleHTTP := httpserver.CreateHTTPServer(httpConfig, "my http server 1")
	simpleHTTP.Start()
}