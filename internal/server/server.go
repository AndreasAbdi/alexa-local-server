package app

import (
	"log"
	"net/http"
	"sync"

	"github.com/AndreasAbdi/alexa-local-server/internal/infrared"

	"github.com/AndreasAbdi/alexa-local-server/internal/cast"
	"github.com/AndreasAbdi/alexa-local-server/internal/config"
	handler "github.com/AndreasAbdi/alexa-local-server/internal/handlers"
	"github.com/AndreasAbdi/alexa-local-server/internal/middleware"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
)

//Server is the instance of the local server to be deployed.
type Server struct {
	do           sync.Once
	conf         config.Wrapper
	castService  *cast.Service
	infraService *infrared.Service
	router       *mux.Router
}

//NewServer is constructor for the server.
func NewServer() *Server {
	conf := config.GetConfig()
	return &Server{
		conf:         conf,
		castService:  cast.NewService(),
		infraService: infrared.NewService(conf),
		router:       mux.NewRouter(),
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
	conf := config.GetConfig()

	s.routes(conf)
	n := negroni.Classic()
	n.UseHandler(s.router)
	err := http.ListenAndServe(s.conf.ServerAddress, n)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

func (s *Server) routes(config config.Wrapper) {

	s.router.HandleFunc("/youtube_test", handler.HandleYoutube(s.castService))
	s.router.HandleFunc("/media_test", handler.HandleMedia(s.castService))
	s.router.HandleFunc("/quit", handler.HandleQuit(s.castService))
	s.router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Test Endpoint.")
	})
	s.router.HandleFunc("/infra", handler.HandleInfra(s.infraService))

	alexaRouter := s.router.PathPrefix("/alexa").Subrouter()
	alexaRouter.HandleFunc("", handler.HandleAlexa(s.conf, s.castService, s.infraService))
	s.router.PathPrefix("/alexa").Handler(negroni.New(
		negroni.HandlerFunc(middleware.GetValidateRequest()),
		negroni.HandlerFunc(middleware.GetVerifyJSON(s.conf.AlexaAppID)),
		negroni.Wrap(alexaRouter),
	))
}
