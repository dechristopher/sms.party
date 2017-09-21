package data

// APIKey is the key used to authenticate requests
type APIKey string

// KeyGenInfo defines the structure for the information required to generate an API key
type KeyGenInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	AppName string `json:"appname"`
	Expires string `json:"expires"`
}

// KeyGenBlob defines the structure for the information stored in redis including api key
type KeyGenBlob struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	AppName string `json:"appname"`
	Expires string `json:"expires"`
	APIKey  APIKey `json:"apikey"`
}

// SMS is the basic struct for messages
type SMS struct {
	Message string `json:"message"`
	Target  string `json:"target"`
}

// CastSMS is the basic struct for sending a single message with multiple targets
type CastSMS struct {
	Targets []string `json:"targets"`
	Message string   `json:"message"`
}

// BatchSMS is the holder for many messages with many targets
type BatchSMS struct {
	Messages []SMS `json:"messages"`
}
