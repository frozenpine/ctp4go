package thost

import "fmt"

type ThostData interface {
	fmt.Stringer

	Type() string
}
