package domain

import "errors"

var ErrSpotNotFound error = errors.New("no spots found")
var ErrUnauthorized error = errors.New("unauthorized")
var ErrInvalidArgument error = errors.New("invalid argument")
var ErrForbidden error = errors.New("forbidden")
