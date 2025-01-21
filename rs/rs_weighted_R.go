package rs

import (
	"container/heap"
	"math/rand"
)

// WeightedReservoirR selects k items from a weighted stream
func WeightedReservoirR[T any](stream []T, weights []float64, k int) []T {
	if len(stream) < k {
		return nil
	}

	reservoir := make([]*Item[T], 0, k)
	pq := PriorityQueue[T]{}
	heap.Init(&pq)

	// Insert first k elements
	for i := 0; i < k; i++ {
		priority := weights[i] / rand.Float64()
		heap.Push(&pq, &Item[T]{stream[i], priority})
	}

	// Process remaining elements
	for i := k; i < len(stream); i++ {
		priority := weights[i] / rand.Float64()
		if pq[0].Priority < priority {
			heap.Pop(&pq)
			heap.Push(&pq, &Item[T]{stream[i], priority})
		}
	}

	// Extract final reservoir
	for _, item := range pq {
		reservoir = append(reservoir, item)
	}

	// Convert back to a slice of T
	result := make([]T, k)
	for i, item := range reservoir {
		result[i] = item.Value
	}
	return result
}
