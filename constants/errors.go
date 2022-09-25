package constants

import "errors"

var (
	ErrServiceIsEmpty = errors.New("Service is empty")
	ErrNameIsEmpty    = errors.New("Name is empty")
	ErrQueueFull      = errors.New("Queue is full (maximum size exceeded)")
)
