package app

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

//Server is the instance of the local server to be deployed.
type Server struct {
	do     sync.Once
	router *mux.Router
}

//NewServer is constructor for the server.
func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

//Init instantiates whatever is needed at the beginning
func (s *Server) Init() {
	s.do.Do(s.onceBody)

}

func (s *Server) onceBody() {
	s.routes()
	http.Handle("/", s.router)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

func (s *Server) routes() {
	s.router.HandleFunc("/youtube_test", s.handleYoutube())
	s.router.HandleFunc("/media_test", s.handleMedia())
	s.router.HandleFunc("/quit", s.handleQuit())
}
