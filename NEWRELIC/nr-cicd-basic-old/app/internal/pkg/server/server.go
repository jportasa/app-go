package server

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
	Shutdown() error
}

type server struct {
	addr string
	srv  *http.Server
}

func (s *server) ListenAndServe() error {
	log.Info("starting server...")
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}

func New(addr string) Server {
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r := mux.NewRouter()
	r.PathPrefix("/").
		Methods("GET").HandlerFunc(helloServer)

	log.Info("Listening at ", addr)
	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(headersOk, originsOk, methodsOk)(r), // Pass our instance of gorilla/mux in.
	}
	return &server{
		addr: addr,
		srv:  srv,
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	log.Info(r.URL)
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
