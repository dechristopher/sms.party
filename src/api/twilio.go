package api

// Credentials for twilio accounts
type Credentials struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

// GetCreds returns a set pf Twilio credentials from Redis
func GetCreds() Credentials {
	var creds Credentials

	// Snag creds from redis here...

	creds.AccountSID = ""
	creds.AuthToken = ""
	creds.FromNumber = ""

	return creds
}
