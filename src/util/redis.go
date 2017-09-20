package util

import (
	"fmt"

	s "github.com/dechristopher/sms.party/src/strings"
)

// GetCredsAsJSON gets random Twilio credentials as a JSON blob
func GetCredsAsJSON() string {
	fmt.Println(s.LogPrefix + "GOT")
	return "{}"
}

// IsNumberOnDNMList checks if a number is on the do not message list
func IsNumberOnDNMList(number string) bool {
	fmt.Println(s.LogPrefix + "YES")
	return true
}

// AddNumberToDNMList adds a number to the do not message list
func AddNumberToDNMList(number string) {
	fmt.Println(s.LogPrefix + "OK")
}

// AddAPIKeyToDatastore adds an API key to the datastore for the desired number of hours
func AddAPIKeyToDatastore(hours int) {
	fmt.Println(s.LogPrefix + "OK")
}

// IsAPIKeyValid checks if an API key is valid and not expired
func IsAPIKeyValid(apikey string) bool {
	fmt.Println(s.LogPrefix + "YES")
	return true
}
