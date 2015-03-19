package mcmq

import (
	"container/list"
)

type MemQueque struct {
	l *list.List
}

func NewMemQueue() *MemQueque {
	return &MemQueque{
		l: list.New(),
	}
}

func (q *MemQueque) Front() (interface{}, error) {
	e := q.l.Front()
	return e, nil
}

func (q *MemQueque) Enque(v interface{}) error {
	q.l.PushBack(v)
	return nil
}

func (q *MemQueque) Deque() (interface{}, error) {
	e := q.l.Front()
	q.l.Remove(e)
	return e, nil
}

func (q *MemQueque) Save() error {
	return nil
}

func (q *MemQueque) Load() error {
	return nil
}
