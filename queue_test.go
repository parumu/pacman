package main

import "testing"

func TestQueue(t *testing.T) {
	q := new(queue)

	q.push(1)
	v := q.pop().(int)
	if v != 1 {
		t.Error("push 1 and pop didn't return 1")
	}

	if q.pop() != nil {
		t.Error("popping empty queue returned something")
	}

	q.push(1)
	q.push(2)
	if q.pop() != 1 || q.pop() != 2 || q.pop() != nil {
		t.Error("2 pushes and pops failed")
	}
}
