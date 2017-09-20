package api

import (
	"fmt"
	"net/http"
	"os"

	d "github.com/dechristopher/sms.party/src/data"
	u "github.com/dechristopher/sms.party/src/util"
	"github.com/sfreiberg/gotwilio"
)

// UnimplementedHandler just Okays everything because it's just there for the ride
func UnimplementedHandler(w http.ResponseWriter, r *http.Request) {
	u.Okay(w)
}

// HostHandler returns container hostname
func HostHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	if name, hnerr := os.Hostname(); hnerr != nil {
		fmt.Printf("Hostname Oopsie: %v %v\n", hnerr, name)
		return
	}
	fmt.Fprintf(w, "%v", name)
}

// IndexHandler serves homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	u.Templates.ExecuteTemplate(w, "index.html", nil)
}

// SendHandler handles basic single SMS sending
func SendHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK if spooled properly
		Otherwise something went wrong
	*/
	// Just for logging yo
	key := d.APIKey(r.Header.Get("apikey"))

	// Snag the deets
	number := r.FormValue("number")
	message := r.FormValue("message")

	// Grab some fancy new creds
	creds := GetCreds()

	// Make a new twilio client using new creds every time
	var twilio = gotwilio.NewTwilioClient(
		creds.AccountSID,
		creds.AuthToken,
	)

	// Send the SMS
	twilio.SendSMS(creds.FromNumber, number, message, "", "")

	// Log it
	u.LogToRedis(r.RemoteAddr+" - "+string(key), r.RequestURI, number, message)

	// Okay it
	u.Okay(w)
}
