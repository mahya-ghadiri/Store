package responses

// Error is the response structure in case of error
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// NewError is the Error factory method
func NewError(status int, message string) *Error {
	return &Error{
		Message: message,
		Status:  status,
	}
}

func (errResponse *Error) Error() string {
	return errResponse.Message
}
