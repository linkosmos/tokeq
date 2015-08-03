package tokeq

import "errors"

// -
var (
	ErrResponseBodyIsEmpty   = errors.New("http.Response Body is empty")
	ErrResponseBodyIsNotHTML = errors.New("http.Response Body is not HTML type or malformed")
)
