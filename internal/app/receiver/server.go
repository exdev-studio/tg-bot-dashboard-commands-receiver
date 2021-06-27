package receiver

import (
	"encoding/json"
	"log"
	"net/http"

	telegram "github.com/exdev-studio/tg-bot-dashboard-commands-receiver/internal/model"
	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleDefaultRequest()).Methods(http.MethodPost)
}

func (s *server) handleDefaultRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		update := &telegram.Update{}
		if err := json.NewDecoder(r.Body).Decode(update); err != nil {
			log.Fatal(err)
			return
		}

		log.Printf("update_id: %d; message: %s; from_id: %d", update.Id, update.Message.Text, update.Message.Chat.Id)

		s.respond(w, http.StatusOK, http.StatusText(http.StatusOK))
	}
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}
