package rs

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestWeightedReservoirSample_Int(t *testing.T) {
	stream := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	weights := []float64{0.1, 0.5, 0.2, 0.3, 0.8, 0.6, 0.9, 0.4, 0.7, 0.2}

	k := 5
	sample := WeightedReservoirR(stream, weights, k)
	if len(sample) != k {
		t.Errorf("Expected sample size %d, got %d", k, len(sample))
	}
}

func TestWeightedReservoirSample_String(t *testing.T) {
	stream := []string{"A", "B", "C", "D", "E", "F", "G"}
	weights := []float64{0.2, 0.5, 0.8, 0.1, 0.7, 0.6, 0.3}

	k := 3
	sample := WeightedReservoirR(stream, weights, k)
	if len(sample) != k {
		t.Errorf("Expected sample size %d, got %d", k, len(sample))
	}
}

func generateTestStreamWR(size int) ([]int, []float64) {
	stream := make([]int, size)
	weights := make([]float64, size)
	for i := range stream {
		stream[i] = i + 1
		weights[i] = rand.Float64() * 10
	}
	return stream, weights
}

func BenchmarkWeightedReservoirSampleR(b *testing.B) {
	sizes := []int{1000, 10000, 100000}
	ks := []int{10, 100, 1000}

	for _, size := range sizes {
		for _, k := range ks {
			b.Run(
				benchmarkName(size, k),
				func(b *testing.B) {
					b.ReportAllocs()

					stream, weights := generateTestStreamWR(size)
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						WeightedReservoirR(stream, weights, k)
					}
				},
			)
		}
	}
}

func benchmarkName(size, k int) string {
	return "StreamSize_" + itoa(size) + "_Reservoir_" + itoa(k)
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
