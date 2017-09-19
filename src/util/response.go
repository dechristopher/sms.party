package util

import (
	"fmt"
	"net/http"

	s "github.com/dechristopher/sms.party/src/strings"
)

// SendResponse sends back a request header and an error string with JSON support
func SendResponse(w http.ResponseWriter, isJSON bool, code int, err string) {
	if !isJSON {
		http.Error(w, err, code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, `%q`, err)
}

// InternalServerError sends a vanilla 500 Internal Server Error
func InternalServerError(w http.ResponseWriter) {
	SendResponse(w, true, 500, `{"error": "`+s.ErrInternalServerError+`"}`)
	return
}

// BadRequest sends a vanilla 400 Bad Request
func BadRequest(w http.ResponseWriter) {
	SendResponse(w, true, 400, `{"error": "`+s.ErrBadRequest+`"}`)
	return
}

// Okay sends 200 Status OK
func Okay(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
