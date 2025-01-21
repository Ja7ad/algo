package rs

import (
	"container/heap"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ReservoirSampleSort selects k elements using a random priority-based min-heap.
func ReservoirSampleSort[T any](stream []T, k int) []T {
	if len(stream) < k {
		return nil
	}

	// Initialize min-heap priority queue
	pq := PriorityQueue[T]{}
	heap.Init(&pq)

	// Process the stream
	for _, item := range stream {
		r := rand.Float64() // Generate random priority between 0 and 1

		if len(pq) < k {
			heap.Push(&pq, &Item[T]{Value: item, Priority: r})
		} else {
			// Keep only k items with largest priority values
			if r > pq[0].Priority {
				heap.Pop(&pq) // Remove the item with the smallest priority
				heap.Push(&pq, &Item[T]{Value: item, Priority: r})
			}
		}
	}

	// Extract final reservoir
	result := make([]T, k)
	for i, item := range pq {
		result[i] = item.Value
	}
	return result
}
