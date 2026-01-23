package errors

import "errors"


var (
	ErrEmptyToken = errors.New("token cannot be empy")
)
