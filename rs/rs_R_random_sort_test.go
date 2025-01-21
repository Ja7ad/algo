package rs

import "testing"

func TestReservoirSampleSort(t *testing.T) {
	N := 1000
	k := 10

	stream := generateTestStreamR(N)
	sample := ReservoirSampleSort(stream, k)

	if len(sample) != k {
		t.Errorf("Expected reservoir size %d, got %d", k, len(sample))
	}

	for _, val := range sample {
		if val < 1 || val > N {
			t.Errorf("Sampled value %d is out of valid range", val)
		}
	}
}

func TestReservoirSampleSort_Probability(t *testing.T) {
	N := 100
	k := 10
	runs := 300000

	stream := generateTestStreamR(N)
	counts := make(map[int]int)

	// Run the algorithm multiple times and count occurrences
	for i := 0; i < runs; i++ {
		sample := ReservoirSampleSort(stream, k)
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

func BenchmarkReservoirSampleSort(b *testing.B) {
	N := 100000
	k := 100
	stream := generateTestStreamR(N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReservoirSampleSort(stream, k)
	}
}
