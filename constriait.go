package ctp4go

import "fmt"

type PtrConstraint[T any] interface {
	*T
}

type DataConstraint interface {
	// fmt.Stringer

	Type() string
}

type DataConstraintPtr[T DataConstraint] interface {
	PtrConstraint[T]

	DataConstraint
}

type Rtn int

func (r Rtn) Error() error {
	if r == 0 {
		return nil
	}

	return fmt.Errorf("%w: code[%d]", ErrRequestFailed, r)
}
