package main

import (
	"yuki.pkg.org/http2server"
)



func main() {
	var server_config = new(http2server.ServerConfig)
	server_config.Port = ":8081"
	server_config.Crt = "https_keys/server.crt"
	server_config.Key = "https_keys/server.key"
	server_config.Static = "static"

	http2server.CreateServer(
		server_config,
	)	
}