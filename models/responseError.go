package models

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

// Implement the error interface
func (e *ResponseError) Error() string {
	return e.Message
}