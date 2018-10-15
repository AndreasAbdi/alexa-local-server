package infrared

import (
	"encoding/json"
	"log"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/imroc/req"
)

const endpointMessage = "msg"
const endpointJSON = "json"
const keyPassword = "pass"
const keyCode = "code"
const keyJSONRaw = "plain"
const codeSeparator = ":"

func sendMessageDefault(url *url.URL, password string, code string, deviceType string, length uint64) {
	request := req.New()
	url.Path = path.Join(url.Path, endpointMessage)
	encodedRequest := strings.Join([]string{code, deviceType, strconv.FormatUint(length, 10)}, codeSeparator)
	params := req.QueryParam{
		keyPassword: password,
		keyCode:     encodedRequest,
	}
	endpoint := url.String()
	log.Print(endpoint)
	_, err := request.Get(url.String(), params)
	if err != nil {
		log.Println("Error with sending request", err)
	}
}

//sends a json request. See https://github.com/mdhiggins/ESP8266-HTTP-IR-Blaster/
func sendMessageJSON(url *url.URL, password string, query Query) {
	request := req.New()
	url.Path = path.Join(url.Path, endpointJSON)
	arrayedQuery := []Query{query}
	jsonRaw, err := json.Marshal(arrayedQuery)
	if err != nil {
		log.Println("Error with converting query into json", err)
	}
	params := req.QueryParam{
		keyPassword: password,
		keyJSONRaw:  string(jsonRaw),
	}
	_, err = request.Post(url.String(), params)
	if err != nil {
		log.Println("Error with sending request", err)
	}
}
