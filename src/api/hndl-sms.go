package api

import (
	"net/http"

	d "github.com/dechristopher/sms.party/src/data"
	u "github.com/dechristopher/sms.party/src/util"
	"github.com/sfreiberg/gotwilio"
)

// SendHandler handles basic single SMS sending
func SendHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK if spooled properly
		400 BR if request body formatted improperly
		500 ISE if something goes wrong
	*/
	// Just for logging yo
	key := d.APIKey(r.Header.Get("apikey"))

	// Snag the deets
	number := r.FormValue("number")
	message := r.FormValue("message")

	// Grab some fancy new creds
	creds, err := GetCreds()

	if err != nil {
		u.InternalServerError(w)
		return
	}

	// Make a new twilio client using new creds every time
	var twilio = gotwilio.NewTwilioClient(
		creds.AccountSID,
		creds.AuthToken,
	)

	// Send the SMS
	twilio.SendSMS(creds.FromNumber, number, message, "", "")

	// Log it
	u.LogToRedis(r.RemoteAddr+" - "+string(key), r.RequestURI, number, message)

	// TODO
	// Update stats for API key and number

	// Okay it
	u.Okay(w)
}
