package strings

const (
	// InfoDefaultPort notifies of default port usage
	InfoDefaultPort = "~ Defaulting port bind to port 80 because no $PORT env var proviced at runtime."
	// ErrBadAPIKey error message for bad API key response
	ErrBadAPIKey = "Invalid cSMS API key"
	// ErrMessageTooLong error message for SMS > 160 characters
	ErrMessageTooLong = "Message too long, max is 160 characters"
	// ErrInsufficientCredit error message for not enough credit on account
	ErrInsufficientCredit = "Insufficient credit on account"
	// ErrInternalServerError error for internal server logic error
	ErrInternalServerError = "A problem has occured within cSMS, please try again momentarily"
	// ErrBadRequest error for invalid request to API
	ErrBadRequest = "Invalid request format, check your formatting"
	// ErrBadPromoCode error for invalid promo code or expired/already used
	ErrBadPromoCode = "Promo code invalid, expired, or already used"
)
