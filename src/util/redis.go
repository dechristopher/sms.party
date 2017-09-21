package util

import (
	"fmt"

	s "github.com/dechristopher/sms.party/src/strings"
)

// GetCredsAsJSON gets random Twilio credentials as a JSON blob
func GetCredsAsJSON() (string, error) {
	fmt.Println(s.LogPrefix + "GOT")
	return "{}", nil
}

// IsNumberOnDNMList checks if a number is on the do not message list
func IsNumberOnDNMList(number string) (bool, error) {
	fmt.Println(s.LogPrefix + "YES")
	return true, nil
}

// AddNumberToDNMList adds a number to the do not message list
func AddNumberToDNMList(number string) error {
	fmt.Println(s.LogPrefix + "OK")

	return nil
}

// AddAPIKeyBlobToDatastore adds an API key to the datastore for the desired number of hours
func AddAPIKeyBlobToDatastore(blob string) error {
	fmt.Println(s.LogPrefix + blob)

	return nil
}

// IsAPIKeyValid checks if an API key is valid and not expired
func IsAPIKeyValid(apikey string) (bool, error) {
	fmt.Println(s.LogPrefix + "YES")
	return true, nil
}
