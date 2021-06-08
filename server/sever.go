package server

import (
	"github.com/dalconoid/url-shortener/storage"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router
}

//New creates new server
func New() *Server {
	s := Server{router: mux.NewRouter()}
	return &s
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

//ConfigureRouter binds handles to routes
func (s *Server) ConfigureRouter(db storage.IURLStorage) {
	s.router.HandleFunc("/alive", handleAlive)
	s.router.HandleFunc("/register-url", handleRegisterShortenedURL(db)).Methods("POST")
	s.router.HandleFunc(`/{slug:[a-zA-Z0-9\-]{3,20}}`, handleRedirectBySlug(db))
}