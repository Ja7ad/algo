package rs

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue_PushPop(t *testing.T) {
	pq := &PriorityQueue[int]{}
	heap.Init(pq)

	items := []struct {
		value    int
		priority float64
	}{
		{1, 0.8},
		{2, 0.5},
		{3, 0.2},
		{4, 0.9},
		{5, 0.1},
	}

	for _, item := range items {
		heap.Push(pq, &Item[int]{Value: item.value, Priority: item.priority})
	}

	expectedOrder := []int{5, 3, 2, 1, 4}
	for i, expected := range expectedOrder {
		popped := heap.Pop(pq).(*Item[int])
		if popped.Value != expected {
			t.Errorf("Expected %d, but got %d at index %d", expected, popped.Value, i)
		}
	}
}

func TestPriorityQueue_GenericTypes(t *testing.T) {
	pq := &PriorityQueue[string]{}
	heap.Init(pq)

	heap.Push(pq, &Item[string]{Value: "apple", Priority: 0.9})
	heap.Push(pq, &Item[string]{Value: "banana", Priority: 0.5})
	heap.Push(pq, &Item[string]{Value: "cherry", Priority: 0.1})

	expectedOrder := []string{"cherry", "banana", "apple"}
	for i, expected := range expectedOrder {
		popped := heap.Pop(pq).(*Item[string])
		if popped.Value != expected {
			t.Errorf("Expected %s, but got %s at index %d", expected, popped.Value, i)
		}
	}
}

func TestPriorityQueue_Length(t *testing.T) {
	pq := &PriorityQueue[float64]{}
	heap.Init(pq)

	heap.Push(pq, &Item[float64]{Value: 1.5, Priority: 0.3})
	heap.Push(pq, &Item[float64]{Value: 2.5, Priority: 0.7})

	if pq.Len() != 2 {
		t.Errorf("Expected length 2, got %d", pq.Len())
	}

	heap.Pop(pq)
	if pq.Len() != 1 {
		t.Errorf("Expected length 1 after pop, got %d", pq.Len())
	}
}

func TestPriorityQueue_Empty(t *testing.T) {
	pq := &PriorityQueue[int]{}
	heap.Init(pq)

	if pq.Len() != 0 {
		t.Errorf("Expected empty queue, got length %d", pq.Len())
	}

	heap.Push(pq, &Item[int]{Value: 10, Priority: 0.5})
	if pq.Len() != 1 {
		t.Errorf("Expected length 1 after push, got %d", pq.Len())
	}

	heap.Pop(pq)
	if pq.Len() != 0 {
		t.Errorf("Expected empty queue after pop, got length %d", pq.Len())
	}
}
