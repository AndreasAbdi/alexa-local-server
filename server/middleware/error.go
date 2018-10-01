package middleware

import (
	"log"
	"net/http"
	"strconv"
)

func httpError(w http.ResponseWriter, logMsg string, err string, errCode int) {
	if logMsg != "" {
		log.Println("[ERROR]: " + "[" + strconv.Itoa(errCode) + "] " + logMsg)
	}

	http.Error(w, err, errCode)
}
