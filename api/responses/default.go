package responses

// Default is the default response structure
type Default struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

// NewDefault is the Default factory method
func NewDefault(status int, message string, data interface{}) *Default {
	return &Default{
		Message: message,
		Data:    data,
		Status:  status,
	}
}

func (response *Default) Error() string {
	return response.Message
}
