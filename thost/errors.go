package thost

import "errors"

var (
	ErrInvalidArgs = errors.New("invalid args")

	ErrLibOpenFailed     = errors.New("lib open failed")
	ErrLibSymbolNotFound = errors.New("lib symbol not found")

	ErrApiCreateFailed = errors.New("api create failed")
	ErrInvalidReqData  = errors.New("invalid request data")
)
