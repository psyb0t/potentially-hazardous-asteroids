package main

import "errors"

var (
	// ErrNASAAPIKeyNotSet is the error returned when the NASA API Key is not set
	ErrNASAAPIKeyNotSet = errors.New("NASA API key not set")
)
