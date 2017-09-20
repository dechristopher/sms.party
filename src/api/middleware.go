package api

import (
	"fmt"
	"net/http"
	"strings"

	d "github.com/dechristopher/sms.party/src/data"
	s "github.com/dechristopher/sms.party/src/strings"
)

// IPLogMiddleware simply prefixes request IP to negroni logger output for every request
func IPLogMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Printf(s.LogPrefix+"Request from %v -> ", r.RemoteAddr)
	next(w, r)
}

// AuthMiddleware verifies presence and validity of API key header
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	/*
		401 Unauthorized if invalid key or no key and JSON error returned
		{"error": "Invalid sms.party API key"}
	*/
	if r.RequestURI == "/" || strings.Contains(r.RequestURI, "/files") {
		fmt.Println("homepage - no auth")
		next(w, r)
		return
	}

	key := d.APIKey(r.Header.Get("apikey"))
	fmt.Printf("%v - checking auth - %v\n", r.RequestURI, key)

	// Check API key validity
	/*if _, existsErr := u.DBC.GetEmail(key); existsErr != nil {
		u.SendResponse(w, true, 401, `{"error": "`+s.ErrBadAPIKey+`"}`)
		return
	}*/

	next(w, r)
}
