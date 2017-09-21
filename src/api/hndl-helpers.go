package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	d "github.com/dechristopher/sms.party/src/data"
	u "github.com/dechristopher/sms.party/src/util"
)

// UnimplementedHandler just Okays everything because it's just there for the ride
func UnimplementedHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK if 1=1
	*/
	u.Okay(w)
}

// KeyStatsHandler gets information and stats about a given API key and returns it as JSON
func KeyStatsHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK if 1=1
	*/
	u.Okay(w)
}

// KeyGenHandler generates a new API key given some info as defined below
func KeyGenHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK if key generated properly
		400 BR if info not provided properly
		500 ISE if something went wrong

		Request Body: application/json in format:
		{
			name: "Andrew DeChristopher", (string)
			email:"drew@kiir.us" (string)
			appname: "cSMS", (string)
			expires: 86400, (int - seconds)
		}
	*/

	// Grab req'd information as JSON blob
	decoder := json.NewDecoder(r.Body)
	var info d.KeyGenInfo
	err := decoder.Decode(&info)
	if err != nil {
		u.SendResponse(w, false, 400, "Invalid key generation information provided. Follow the JSON format in the docs.")
		return
	}
	defer r.Body.Close()

	// Generate API key
	apikey := d.GenAPIKey()

	// Generate a storage blob
	blob := &d.KeyGenBlob{
		Name:    info.Name,
		Email:   info.Email,
		AppName: info.AppName,
		Expires: info.Expires,
		APIKey:  apikey,
	}

	// Marshall it as JSON
	jsonblob, err := json.Marshal(blob)
	if err != nil {
		fmt.Println(err)
		u.InternalServerError(w)
		return
	}

	// Store it in Redis for future validation and retreival
	if err := u.AddAPIKeyBlobToDatastore(string(jsonblob)); err != nil {
		u.InternalServerError(w)
		return
	}

	// Success
	u.Okay(w)
}

// HostHandler returns container hostname
func HostHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK with hostname in response body
		500 ISE if hostname cannot be resolved
	*/
	var name string
	if name, hnerr := os.Hostname(); hnerr != nil {
		fmt.Printf("Hostname Oopsie: %v %v\n", hnerr, name)
		u.InternalServerError(w)
		return
	}
	fmt.Fprintf(w, "%v", name)
}
