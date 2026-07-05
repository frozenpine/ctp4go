package thost

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCreator  = errors.New("invalid creator")
	ErrCreatorMissing  = errors.New("thost api creator missing")
	ErrCreatorConflict = errors.New(
		"multiple version's thost api creator conflict",
	)

	ErrInvalidArgs = errors.New("invalid args")

	ErrLibOpenFailed     = errors.New("lib open failed")
	ErrLibSymbolNotFound = errors.New("lib symbol not found")

	ErrApiCreateFailed = errors.New("api create failed")
	ErrInvalidReqData  = errors.New("invalid request data")
	ErrRequestFailed   = errors.New("request failed")
)

type Rtn struct {
	Code int
}

func (r Rtn) Error() error {
	if r.Code == 0 {
		return nil
	}

	return fmt.Errorf("%w: code[%d]", ErrRequestFailed, r.Code)
}
