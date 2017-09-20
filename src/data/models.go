package data

// APIKey is the key used to authenticate requests
type APIKey string

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
