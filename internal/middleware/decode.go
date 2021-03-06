package middleware

import (
	"net/http"

	"github.com/AndreasAbdi/alexa-local-server/internal/encoding"
)

//GetVerifyJSON Decode the JSON request and verify it.
func GetVerifyJSON(appID string) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := r.Context()
		echoReq, err := encoding.GetRequest(ctx, r)
		if err != nil {
			httpError(w, err.Error(), "Bad Request", 400)
			return
		}
		// Check the timestamp
		if !echoReq.VerifyTimestamp() && r.URL.Query().Get("_dev") == "" {
			httpError(w, "Request too old to continue (>150s).", "Bad Request", 400)
			return
		}

		// Check the app id
		if !echoReq.VerifyAppID(appID) {
			httpError(w, "Echo AppID mismatch!", "Bad Request", 400)
			return
		}
		next(w, r)
	}
}
