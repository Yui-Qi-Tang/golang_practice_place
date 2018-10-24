package httpserver

// HTTPSConfig : config for https setup
type HTTPSConfig struct {
	Port           string
	Crt            string
	Key            string
}

// HTTPConfig : config for http setup
type HTTPConfig struct {
	Port           string
}
