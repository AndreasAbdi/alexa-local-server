package als

import (
	"fmt"
	"sync"

	"github.com/gorilla/mux"
)

type server struct {
	do     sync.Once
	router mux.Router
}

func (s *server) Init() {
	s.do.Do(s.DoInit)
}

func (s *server) DoInit() {
	fmt.Print("Hello world from server init")
}
