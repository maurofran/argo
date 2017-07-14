package argo

// IllegalArgumentError is the error returned when an argument is invalid.
type IllegalArgumentError struct {
	message string
}

func (err *IllegalArgumentError) Error() string {
	return err.message
}

// IllegalStateError is the error returned when the state of an object is invalid.
type IllegalStateError struct {
	message string
}

func (err *IllegalStateError) Error() string {
	return err.message
}
