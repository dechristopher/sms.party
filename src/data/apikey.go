package data

import "math/rand"

// GenAPIKey generates an API key following standard length and complexity requirements
func GenAPIKey() APIKey {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%")

	rnd := make([]rune, 9)

	for i := range rnd {
		rnd[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return APIKey(rnd)
}
