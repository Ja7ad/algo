package rs

import (
	"fmt"
	"testing"
)

func generateTestStreamR(N int) []int {
	stream := make([]int, N)
	for i := range stream {
		stream[i] = i + 1
	}
	return stream
}

func TestReservoirSampleR(t *testing.T) {
	N := 1000
	k := 10

	stream := generateTestStreamR(N)
	sample := ReservoirSampleR(stream, k)

	if len(sample) != k {
		t.Errorf("Expected reservoir size %d, got %d", k, len(sample))
	}

	for _, val := range sample {
		if val < 1 || val > N {
			t.Errorf("Sampled value %d is out of valid range", val)
		}
	}
}

func TestReservoirSampleR_Probability(t *testing.T) {
	N := 100
	k := 10
	runs := 300000

	stream := generateTestStreamR(N)
	counts := make(map[int]int)

	// Run the algorithm multiple times and count occurrences
	for i := 0; i < runs; i++ {
		sample := ReservoirSampleR(stream, k)
		for _, val := range sample {
			counts[val]++
		}
	}

	// Expected probability
	expected := float64(runs) * float64(k) / float64(N)
	tolerance := expected * 0.05 // Allow ±5% error

	for i := 1; i <= N; i++ {
		if count, exists := counts[i]; exists {
			if float64(count) < expected-tolerance || float64(count) > expected+tolerance {
				t.Errorf("Item %d has count %d, expected ~%f", i, count, expected)
			}
		}
	}
}

func BenchmarkReservoirSampleR(b *testing.B) {
	sizes := []int{1000, 10000, 100000}
	ks := []int{10, 100, 1000}

	for _, size := range sizes {
		for _, k := range ks {
			b.Run(
				fmt.Sprintf("StreamSize_%d_Reservoir_%d", size, k),
				func(b *testing.B) {
					b.ReportAllocs()
					stream := generateTestStreamR(size)
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						ReservoirSampleR(stream, k)
					}
				},
			)
		}
	}
}
