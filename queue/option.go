package queue

type OptionFunc func(*Queue) *Queue

func WithName(name string) OptionFunc {
	return func(q *Queue) *Queue {
		q.Name = name
		return q
	}
}

func WithMaxSize(max uint) OptionFunc {
	return func(q *Queue) *Queue {
		q.maxSize = max
		return q
	}
}
