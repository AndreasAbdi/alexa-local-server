package app

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/AndreasAbdi/alexa-local-server/server/encoding"
	"github.com/AndreasAbdi/alexa-local-server/server/middleware"
	"github.com/mikeflynn/go-alexa/skillserver"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
)

//Server is the instance of the local server to be deployed.
type Server struct {
	do              sync.Once
	appID           string
	address         string
	encodingService *encoding.Service
	router          *mux.Router
}

//NewServer is constructor for the server.
func NewServer(address string, appID string) *Server {
	return &Server{
		address:         address,
		appID:           appID,
		encodingService: &encoding.Service{},
		router:          mux.NewRouter(),
	}
}

//Init instantiates whatever is needed at the beginning
func (s *Server) Init() {
	s.do.Do(s.onceBody)
}

//ServeHTTP passes request to the server going through all the routes and middlewares.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) onceBody() {
	s.routes()
	alexaRouter := s.router.PathPrefix("/alexa").Subrouter()
	alexaRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		alexaResp := skillserver.NewEchoResponse()
		json, _ := alexaResp.String()
		fmt.Print("Got alexa request")
		log.Print("Got an alexa request in log!")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(json)
	})
	s.router.PathPrefix("/alexa").Handler(negroni.New(
		negroni.HandlerFunc(middleware.GetValidateRequest()),
		negroni.HandlerFunc(middleware.GetVerifyJSON(s.appID, s.encodingService)),
		negroni.Wrap(alexaRouter),
	))
	n := negroni.Classic()
	n.UseHandler(s.router)
	err := http.ListenAndServe(s.address, n)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

func (s *Server) routes() {

	s.router.HandleFunc("/youtube_test", s.handleYoutube())
	s.router.HandleFunc("/media_test", s.handleMedia())
	s.router.HandleFunc("/quit", s.handleQuit())
	s.router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Test Endpoint.")
	})
}
