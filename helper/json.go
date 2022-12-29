package helper

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestBody(request *http.Request, body interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(body)
	PanicIfError(err)
}

func EncodeResponseBody(writer http.ResponseWriter, resBody interface{}) {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(resBody)
	PanicIfError(err)
}
