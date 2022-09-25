package constants

import "errors"

var (
	ErrQueueFull = errors.New("Queue is full (maximum size exceeded)")
)
