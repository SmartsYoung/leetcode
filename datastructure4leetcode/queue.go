package main

import (
	"container/list"
	"sync"
)

type Queue struct {
	l *list.List
	m sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (q *Queue) Push(v interface{}) {
	if v == nil {
		return
	}
	q.m.Lock()
	defer q.m.Unlock()
	q.l.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	q.m.Lock()
	defer q.m.Unlock()
	if elem := q.l.Front(); elem != nil {
		q.l.Remove(elem)
		return elem.Value
	}
	return nil
}

func (q *Queue) Len() int {
	q.m.Lock()
	defer q.m.Unlock()
	return q.l.Len()
}
