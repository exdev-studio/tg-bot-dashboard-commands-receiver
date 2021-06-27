package receiver

import (
	"log"
	"net/http"
)

func Start(config *Config) error {
	log.Printf("server starting on %s addr", config.BindAddr)
	s := newServer()

	return http.ListenAndServe(config.BindAddr, s)
}
