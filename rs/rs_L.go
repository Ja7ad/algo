package rs

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ReservoirSampleL selects k elements from a stream using Algorithm L
func ReservoirSampleL[T any](stream []T, k int) []T {
	if len(stream) < k {
		return nil // Not enough elements
	}

	// Step 1: Fill the initial reservoir
	reservoir := make([]T, k)
	copy(reservoir, stream[:k])

	// Step 2: Initialize weight factor W
	W := math.Exp(math.Log(rand.Float64()) / float64(k))

	i := k // Current position in the stream

	// Step 3: Process remaining elements with skipping
	for i < len(stream) {
		// Calculate number of elements to skip
		skip := int(math.Floor(math.Log(rand.Float64()) / math.Log(1-W)))
		i += skip + 1 // Move forward in the stream

		// If within bounds, replace a random item in the reservoir
		if i < len(stream) {
			j := rand.Intn(k) // Random index in the reservoir
			reservoir[j] = stream[i]

			// Update weight factor W
			W *= math.Exp(math.Log(rand.Float64()) / float64(k))
		}
	}

	return reservoir
}
