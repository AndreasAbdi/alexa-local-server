package alexa

//RequestTypeError where the request type file in the alexa request was not valid.
type RequestTypeError struct {
	RequestType string
}

func (e *RequestTypeError) Error() string {
	return "Invalid request type: " + e.RequestType
}
