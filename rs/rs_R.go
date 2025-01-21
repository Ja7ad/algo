package rs

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ReservoirSampleR selects k elements from a stream of unknown size using Algorithm R.
func ReservoirSampleR[T any](stream []T, k int) []T {
	if len(stream) < k {
		return nil // Not enough elements
	}

	// Step 1: Fill the reservoir with the first k elements
	reservoir := make([]T, k)
	copy(reservoir, stream[:k])

	// Step 2: Process remaining elements with decreasing probability
	for i := k; i < len(stream); i++ {
		// Select a random index in range [0, i]
		j := rand.Intn(i + 1)

		// If the random index falls within the reservoir size, replace the element
		if j < k {
			reservoir[j] = stream[i]
		}
	}

	return reservoir
}
