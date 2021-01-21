package utils

type errorMessage struct {
	message string
}

// NewError used to create the error message and return the error
func NewError(message string) error {
	return &errorMessage{message}
}

func (e *errorMessage) Error() string {
	return e.message
}
