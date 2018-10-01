package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mikeflynn/go-alexa/skillserver"
)

func (s *Server) handleAlexa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alexaResp := skillserver.NewEchoResponse()
		json, _ := alexaResp.String()
		fmt.Print("Got alexa request")
		log.Print("Got an alexa request in log!")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(json)
	}
}
