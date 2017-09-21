package api

import (
	"fmt"

	u "github.com/dechristopher/sms.party/src/util"
)

// Credentials for twilio accounts
type Credentials struct {
	AccountSID string `json:"sid"`
	AuthToken  string `json:"auth"`
	FromNumber string `json:"num"`
}

// GetCreds returns a set pf Twilio credentials from Redis
func GetCreds() (Credentials, error) {
	var creds Credentials

	// Snag creds from redis here...
	blob, err := u.GetCredsAsJSON()

	fmt.Println(blob)

	creds.AccountSID = ""
	creds.AuthToken = ""
	creds.FromNumber = ""

	return creds, err
}
