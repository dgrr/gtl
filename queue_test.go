package gtl

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var q Queue[int]

	if q.Pop().HasValue() {
		t.Fatal("first pop is not nil?")
	}

	q.PushBack(1)
	if e := q.Pop(); !e.HasValue() || e.Get() != 1 {
		t.Fatal("Unexpected pop element")
	}

	if q.Pop().HasValue() {
		t.Fatal("pop after removed element")
	}

	for i := 2; i < 20; i++ {
		q.PushBack(i)
	}

	q.PushFront(1)
	q.PushFront(0)

	for i := 0; i < 20; i++ {
		if e := q.Pop(); !e.HasValue() || e.Get() != i {
			t.Fatalf("Unexpected: %d <> %d", e.Get(), i)
		}
	}

	if q.Pop().HasValue() {
		t.Fatal("pop after removed element")
	}
}
