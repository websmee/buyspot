package domain

import "errors"

var ErrSpotNotFound error = errors.New("no spots found")
var ErrCurrentPricesNotFound error = errors.New("current prices not found")
var ErrUnauthorized error = errors.New("unauthorized")
var ErrInvalidArgument error = errors.New("invalid argument")
