package queue

import (
	"testing"

	"github.com/golang-must/must"
	"github.com/hadihammurabi/sungoq/constants"
)

type TestData struct {
	ID int
}

var testDatum []TestData = []TestData{
	{1},
	{2},
	{3},
	{4},
	{5},
}

func TestNewQueue(t *testing.T) {

	must := must.New(t)

	q := New()

	must.Equal(q.Name, "")
	must.Equal(len(q.storage), 0)
}

func TestNewQueueWithName(t *testing.T) {

	must := must.New(t)

	name := "zzzzz"

	q := New(
		WithName(name),
	)

	must.Equal(q.Name, name)
}

func TestNewQueueWithMaxSize(t *testing.T) {

	must := must.New(t)

	maxSize := uint(2)

	q := New(
		WithMaxSize(maxSize),
	)

	must.Equal(q.maxSize, maxSize)
}

func TestEnqueue(t *testing.T) {

	must := must.New(t)

	q := New()

	for _, data := range testDatum {
		err := q.Enqueue(data)

		// check new data position is in the last of storage
		must.Equal(data.ID, q.storage[q.size-1].(TestData).ID)
		must.Nil(err)
	}

	must.Equal(uint(len(testDatum)), q.size)
}

func TestEnqueueWithMaxSize(t *testing.T) {

	must := must.New(t)

	maxSize := uint(2)

	q := New(
		WithMaxSize(maxSize),
	)

	for i, data := range testDatum {
		err := q.Enqueue(data)

		if i > int(q.maxSize) {
			must.NotNil(err)
			must.Equal(err, constants.ErrQueueFull)
		}
	}

}

func TestDequeue(t *testing.T) {

	must := must.New(t)

	q := New()

	for _, data := range testDatum {
		q.Enqueue(data)
	}

	for _, data := range testDatum {
		dataFromQ := q.Dequeue()
		dataFromQConverted := dataFromQ.(TestData)

		must.Equal(data.ID, dataFromQConverted.ID)
	}

	must.Equal(q.size, uint(0))
}
