package httpserver

type HttpsConfig struct {
	Port string
	Crt string
	Key string
	StaticFilePath string
}

type HttpConfig struct {
	Port string
	StaticFilePath string
}

