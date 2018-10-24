package main

import 	"yuki.pkg.org/httpserver"
import "fmt"

func test(w interface{}, r interface{}) {
	fmt.Println(w,r)
}

func main() {
	// Pass test function to SetHandleFunc
	httpserver.SetHandleFunc(test)
	httpConfig := &httpserver.HTTPConfig{Port: ":8001", StaticFilePath: "./test"}
	simpleHTTP := httpserver.CreateHTTPServer(httpConfig, "my http server 1")
	simpleHTTP.Start()
}