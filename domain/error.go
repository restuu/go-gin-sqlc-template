package domain

import "errors"

// collection of errors
var (
	ErrNoData              = errors.New("no_data")
	ErrInternalServerError = errors.New("internal_server_error")
)
