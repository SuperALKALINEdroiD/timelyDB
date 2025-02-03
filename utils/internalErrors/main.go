package internalerrors

import "errors"

var FileNotFound = errors.New("Requested file not found.")
var InvalidPath = errors.New("invalid storage path")
