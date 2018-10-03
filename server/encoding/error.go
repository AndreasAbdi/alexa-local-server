package encoding

//Error describing encoding/unmarshalling failures.
type Error struct {
	Internal error
}

func (e *Error) Error() string {
	return "Failed to decode " + e.Internal.Error()
}
