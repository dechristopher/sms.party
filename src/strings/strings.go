package strings

const (
	// LogPrefix is exactly what it sounds like
	LogPrefix = "[sms.p] "
	// InfoStartup notifies of service startup
	InfoStartup = "~ Init sms.party backend v"
	// InfoDefaultPort notifies of default port usage
	InfoDefaultPort = "~ Defaulting port bind to port 80 because no $PORT env var proviced at runtime."
	// ErrBadAPIKey error message for bad API key response
	ErrBadAPIKey = "Invalid sms.party API key"
	// ErrMessageTooLong error message for SMS > 160 characters
	ErrMessageTooLong = "Message too long, max is 160 characters"
	// ErrInternalServerError error for internal server logic error
	ErrInternalServerError = "A problem has occured within sms.party, please try again momentarily"
	// ErrBadRequest error for invalid request to API
	ErrBadRequest = "Invalid request format, check your formatting"
	// ErrRedisConnection
	ErrRedisConnection = "Redis server could not be contacted, service shutting down."
)
