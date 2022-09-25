package queue

import (
	"github.com/hadihammurabi/sungoq/constants"
)

type Queue struct {
	Name    string
	storage []interface{}

	maxSize uint
	size    uint
}

func New(options ...OptionFunc) *Queue {
	q := &Queue{}

	for _, opt := range options {
		opt(q)
	}

	return q
}

func (q *Queue) Enqueue(data interface{}) error {
	if q.maxSize > 0 && q.size == q.maxSize {
		return constants.ErrQueueFull
	}

	q.storage = append(q.storage, data)
	q.size++

	return nil
}

func (q *Queue) Dequeue() interface{} {
	data := q.storage[0]
	q.storage = append(q.storage[:0], q.storage[1:]...)
	q.size--

	return data
}
