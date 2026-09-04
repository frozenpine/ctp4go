package ctp4go

import "errors"

// lib initialize errors
var (
	ErrInvalidArgs       = errors.New("invalid args")
	ErrLibOpenFailed     = errors.New("lib open failed")
	ErrLibSymbolNotFound = errors.New("lib symbol not found")
	ErrApiCreateFailed   = errors.New("api create failed")
)

// ffi errors
var (
	ErrVecIdxOutOfRange = errors.New("index out of range")
)

// api request errors
var (
	ErrInvalidState  = errors.New("invalid state")
	ErrRequestFailed = errors.New("request failed")
)
