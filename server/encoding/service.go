package encoding

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mikeflynn/go-alexa/skillserver"
)

//Service to grab an encoded version of an echorequest.
type Service struct {
}

//GetRequest attempts to retrieve an unmarshalled value of an alexa request.
func (s *Service) GetRequest(ctx context.Context, r *http.Request) (*skillserver.EchoRequest, error) {
	var echoReq *skillserver.EchoRequest
	err := json.NewDecoder(r.Body).Decode(&echoReq)
	return echoReq, err
}
