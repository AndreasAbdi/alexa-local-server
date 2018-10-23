package alexa

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/AndreasAbdi/alexa-local-server/internal/encoding"
	"github.com/mikeflynn/go-alexa/skillserver"
)

const launchRequestType = "LaunchRequest"
const intentRequestType = "IntentRequest"
const sessionEndedRequestType = "sessionEndedRequest"
const audioPlayerRequestTypePrefix = "AudioPlayer."

var validRequestTypes = map[string]bool{
	launchRequestType:       true,
	intentRequestType:       true,
	sessionEndedRequestType: true,
}

//HandleAlexaRequest for dealing with alexa requests
func HandleAlexaRequest(app App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Got an alexa request in log!")
		ctx := r.Context()
		alexaReq, err := encoding.GetRequest(ctx, r)
		if err != nil {
			http.Error(w, (&encoding.Error{Internal: err}).Error(), http.StatusBadRequest)
			return
		}
		if !hasValidRequestType(ctx, alexaReq) {
			http.Error(w, (&RequestTypeError{alexaReq.GetRequestType()}).Error(), http.StatusBadRequest)
			return
		}
		runProperCommand(ctx, app, w, alexaReq)
	}
}

func hasValidRequestType(ctx context.Context, r *skillserver.EchoRequest) bool {
	requestType := r.GetRequestType()
	_, ok := validRequestTypes[requestType]
	ok = ok || strings.HasPrefix(requestType, audioPlayerRequestTypePrefix)
	return ok
}

func runProperCommand(ctx context.Context, app App, w http.ResponseWriter, req *skillserver.EchoRequest) {
	switch req.GetRequestType() {
	case launchRequestType:
		if app.LaunchHandler == nil {
			return
		}
		app.LaunchHandler(ctx, w, req)
	case intentRequestType:
		if app.IntentHandler == nil {
			return
		}
		app.IntentHandler(ctx, w, req)
	case sessionEndedRequestType:
		if app.SessionEndedHandler == nil {
			return
		}
		app.SessionEndedHandler(ctx, w, req)
	default:
		if strings.HasPrefix(launchRequestType, audioPlayerRequestTypePrefix) {
			if app.AudioPlayerStateHandler == nil {
				return
			}
			app.AudioPlayerStateHandler(ctx, w, req)
		}
	}
}

//HandlerFunc is function for alexa response/request objects
type HandlerFunc func(context.Context, http.ResponseWriter, *skillserver.EchoRequest)

//App is a mapping of the functions and intents for an alexa app.
type App struct {
	AppID                   string
	LaunchHandler           HandlerFunc
	IntentHandler           HandlerFunc
	SessionEndedHandler     HandlerFunc
	AudioPlayerStateHandler HandlerFunc
}
