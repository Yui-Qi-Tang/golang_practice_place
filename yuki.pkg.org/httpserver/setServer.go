package httpserver

import (
	"net/http"

	"yuki.pkg.org/tools/logger"
)

/*
type tlsInfo interface {
	certFile() string
	keyFile() string
}
*/

type server struct {
	name     string
	tls      bool // consider map? {} / {certFile: '/path/to/file', keyFile: '/p/t/f'}
	certFile string
	keyFile  string
	instance *http.Server
}

// CreateHTTPServer : create server support Http
func CreateHTTPServer(c *HTTPConfig, serverName string) *server {
	// set Http Server with c
	logTag := "CreateHTTPServer()"

	logger.InfoLog("Create net/http server", logTag)
	s := &http.Server{
		Addr: c.Port,
		// Handler:        myHandler,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	logger.InfoLog("Create server instance", logTag)
	newServer := new(server)
	newServer.name = serverName
	newServer.tls = false
	newServer.instance = s

	return newServer
}

// CreateHTTPTlsServer : create server support Https
func CreateHTTPTlsServer(c *HTTPSConfig, serverName string) *server {
	logTag := "CreateHTTPTlsServer()"
	logger.InfoLog("Create net/http Tls server", logTag)
	s := &http.Server{
		Addr: c.Port,
		// Handler:        myHandler,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	logger.InfoLog("Create tls server instance", logTag)
	newServer := new(server)
	newServer.name = serverName
	newServer.tls = true
	newServer.certFile = c.Crt
	newServer.keyFile = c.Key
	newServer.instance = s

	return newServer
}

// Start : start service
func (s *server) Start() {
	logger.InfoLog("Server: "+s.name+" is starting at "+s.instance.Addr+"...", "Start()")

	if !s.tls {
		logger.InfoLog("Http server Listening...", "Start()")
		s.instance.ListenAndServe()
	} else {
		logger.InfoLog("Https server Listening...", "Start()")
		s.instance.ListenAndServeTLS(s.certFile, s.keyFile)
	}
}

// StartRutine : wrap Start() for gorutine
func (s *server) StartRutine(ch chan bool) {
	logger.InfoLog("Start http service with gorutine", "StartRutine()")
	s.Start()
	ch <- false
}
