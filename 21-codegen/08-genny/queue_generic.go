package queue

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Generic=string,int"

type Generic generic.Type

// GenericQueue represents a queue of Generic types.
type GenericQueue struct {
	items []Generic
}

// NewGenericQueue makes a new empty Generic queue.
func NewGenericQueue() *GenericQueue {
	return &GenericQueue{items: make([]Generic, 0)}
}

// Enq adds an item to the queue.
func (q *GenericQueue) Enq(obj Generic) *GenericQueue {
	q.items = append(q.items, obj)
	return q
}

// Deq removes and returns the next item in the queue.
func (q *GenericQueue) Deq() Generic {
	obj := q.items[0]
	q.items = q.items[1:]
	return obj
}

// Len gets the current number of Generic items in the queue.
func (q *GenericQueue) Len() int {
	return len(q.items)
}

func minGeneric(v1, v2 Generic) Generic {
	if v1 > v2 {
		return v2
	}
	return v1
}
