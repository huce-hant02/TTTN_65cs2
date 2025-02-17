package error

// Common error codes
const (
	// ErrorInternalServerError is a constant for error code when internal server error
	ErrorInternalServerError = "INTERNAL_SERVER_ERROR"
)

// User error codes
const (
	// ErrorCodeInvalidUsernameOrPassword is a constant for error code when username or password is invalid
	ErrorCodeInvalidUsernameOrPassword = "INVALID_USERNAME_OR_PASSWORD"
	// ErrorCodeUserNotFound is a constant for error code when user is not found
	ErrorCodeUserNotFound = "USER_NOT_FOUND"
	// ErrorCodeUserAlreadyExists is a constant for error code when user already exists
	ErrorCodeUserAlreadyExists = "USER_ALREADY_EXISTS"
	// ErrorCodeInvalidToken is a constant for error code when token is invalid
	ErrorCodeInvalidToken = "INVALID_TOKEN"
)
