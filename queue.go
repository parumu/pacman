package main

type queueNd struct {
	v    interface{}
	next *queueNd
}

type queue struct {
	first *queueNd
	last  *queueNd
}

func (q *queue) push(v interface{}) {
	nd := &queueNd{v, nil}

	if q.first == nil { // if empty
		q.first = nd
		q.last = nd
	} else if q.first == q.last { // if 1 element
		q.first.next = nd
		q.last = nd
	} else { // if 2 or more elements
		q.last.next = nd
		q.last = nd
	}
}

func (q *queue) pop() interface{} {
	if q.first == nil { // if empty
		return nil
	}
	if q.first == q.last { // if 1 element
		v := q.first.v
		q.first = nil
		q.last = nil
		return v
	}
	// if 2 or more elements
	v := q.first.v
	q.first = q.first.next
	return v
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}
