package datahandlers

import "errors"

var (
	// ErrStatusNotOK is the error returned when an HTTP response status is not OK
	ErrStatusNotOK = errors.New("response status not OK")

	// ErrCacheItemDataUnexpectedDataType is the error response returned when
	// a cached item's data if of an unexpected data type
	ErrCacheItemDataUnexpectedDataType = errors.New("cache item data is of unexpected data type")
)
