package rws

import (
	"math/rand"
	"sort"
)

// WeightedSelector represents a structure for weighted random selection with binary search optimization.
type WeightedSelector[T any] struct {
	items         []T
	cumulativeSum []int
	total         int
}

// NewWeightedSelector creates a selector with explicit weights and prepares cumulative sums.
func NewWeightedSelector[T any](weightedItems map[int]T) (*WeightedSelector[T], error) {
	if len(weightedItems) == 0 {
		return nil, ErrEmptyMapItems
	}

	var items []T
	var cumulativeSum []int
	total := 0

	for weight, item := range weightedItems {
		if weight <= 0 {
			return nil, ErrInvalidWeight
		}
		items = append(items, item)
		total += weight
		cumulativeSum = append(cumulativeSum, total) // Store cumulative weights
	}

	return &WeightedSelector[T]{items, cumulativeSum, total}, nil
}

// NewAutoWeightedSelector assigns random weights to items and computes cumulative sums.
func NewAutoWeightedSelector[T any](items []T) (*WeightedSelector[T], error) {
	if len(items) == 0 {
		return nil, ErrEmptyItems
	}

	rand.Seed(rand.Int63()) // Seed randomness
	var cumulativeSum []int
	total := 0

	weights := make([]int, len(items))
	for i := range items {
		weights[i] = rand.Intn(100) + 1 // Random weight 1-100
		total += weights[i]
		cumulativeSum = append(cumulativeSum, total)
	}

	return &WeightedSelector[T]{items, cumulativeSum, total}, nil
}

// Pick selects an item on cumulative weights.
func (ws *WeightedSelector[T]) Pick() (T, error) {
	if len(ws.items) == 0 {
		var zeroValue T
		return zeroValue, ErrNullItems
	}

	// Generate a random number between 0 and total weight (exclusive)
	r := rand.Intn(ws.total)

	// Perform binary search to find the correct item
	idx := sort.Search(len(ws.cumulativeSum), func(i int) bool {
		return ws.cumulativeSum[i] > r
	})

	return ws.items[idx], nil
}
