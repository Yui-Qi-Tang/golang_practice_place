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

// Exported those data type to set inject function for SetHandleFun use
type (
	Resp http.ResponseWriter
	Req *http.Request
)

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

// setHandleFunc : return func(w http.ResponseWriter, r *http.Request)
func setHandleFunc(inject func(w Resp, r Req)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		inject(w, r)
	}
}
// SetHandleFunc : A setHandlerFunc wrapper
func SetHandleFunc(path string, inject func(w Resp, r Req)) {
	logger.InfoLog("Set Handle function", "SetHandleFunc")
	s := setHandleFunc(inject)
	http.HandleFunc(path, s)
}

// SetStaticFile
func SetStaticFile(staticUrl string, staticFilePath string) {
	http.Handle(staticUrl, http.StripPrefix(staticUrl, http.FileServer(http.Dir(staticFilePath))))
}