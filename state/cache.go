package state

import "context"

type Cache[T any] struct {
	ctx    context.Context
	cancel context.CancelFunc

	name  string
	data  []T
	count int
}
