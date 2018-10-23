package encoding

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mikeflynn/go-alexa/skillserver"
)

//GetRequest attempts to retrieve an unmarshalled value of an alexa request.
func GetRequest(ctx context.Context, r *http.Request) (*skillserver.EchoRequest, error) {
	var echoReq *skillserver.EchoRequest
	err := json.NewDecoder(r.Body).Decode(&echoReq)
	return echoReq, err
}
