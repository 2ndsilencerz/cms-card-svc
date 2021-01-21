package utils

type errorMessage struct {
	message string
}

// New used to create the error message and return the error
func New(message string) error {
	return &errorMessage{message}
}

func (e *errorMessage) Error() string {
	return e.message
}
