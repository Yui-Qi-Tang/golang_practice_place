package httpserver

import "yuki.pkg.org/tools/logger"

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
	ch <- false // TODO: add set ch as flase condition
}

/*
func (s *server) setStaticFile() {
	s.instance.Handle(
		"/static/",
		s.instance.StripPrefix("/static/", s.instance.FileServer(s.instance.Dir(s.StaticFilePath))),
	)
}
*/