package domain

import "errors"

var ErrSpotNotFound = errors.New("no spots found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrInvalidArgument = errors.New("invalid argument")
var ErrForbidden = errors.New("forbidden")
