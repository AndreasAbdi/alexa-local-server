package alexa

import (
	"net/http"

	"github.com/mikeflynn/go-alexa/skillserver"
)

//WriteResponse to the responsewriter
func WriteResponse(w http.ResponseWriter, resp *skillserver.EchoResponse) {
	json, _ := resp.String()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Write(json)
}
