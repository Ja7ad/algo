package rws

import "testing"

func TestNewWeightedSelector(t *testing.T) {
	tests := []struct {
		name          string
		weightedItems map[int]string
		expectErr     error
	}{
		{"Valid selection", map[int]string{3: "Apple", 1: "Banana", 6: "Cherry"}, nil},
		{"Empty map", map[int]string{}, ErrEmptyMapItems},
		{"Negative weight", map[int]string{-2: "Invalid", 3: "Apple"}, ErrInvalidWeight},
		{"Zero weight", map[int]string{0: "Zero", 3: "Apple"}, ErrInvalidWeight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector, err := NewWeightedSelector(tt.weightedItems)
			if err != nil && err != tt.expectErr {
				t.Errorf("Expected error %v, got %v", tt.expectErr, err)
			}
			if err == nil && selector == nil {
				t.Errorf("Selector should not be nil for valid input")
			}
		})
	}
}

func TestNewAutoWeightedSelector(t *testing.T) {
	tests := []struct {
		name      string
		items     []string
		expectErr error
	}{
		{"Valid selection", []string{"Dog", "Cat", "Fish"}, nil},
		{"Empty slice", []string{}, ErrEmptyItems},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector, err := NewAutoWeightedSelector(tt.items)
			if err != nil && err != tt.expectErr {
				t.Errorf("Expected error %v, got %v", tt.expectErr, err)
			}
			if err == nil && selector == nil {
				t.Errorf("Selector should not be nil for valid input")
			}
		})
	}
}

func TestPick(t *testing.T) {
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	selector, err := NewWeightedSelector(weightedItems)
	if err != nil {
		t.Fatalf("Failed to create selector: %v", err)
	}

	for i := 0; i < 100; i++ {
		item, err := selector.Pick()
		if err != nil {
			t.Errorf("Pick() returned an error: %v", err)
		}
		if item != "Apple" && item != "Banana" && item != "Cherry" {
			t.Errorf("Pick() returned an unexpected item: %v", item)
		}
	}
}

func TestProbabilityDistribution(t *testing.T) {
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	selector, err := NewWeightedSelector(weightedItems)
	if err != nil {
		t.Fatalf("Failed to create selector: %v", err)
	}

	trials := 100000
	counts := map[string]int{}

	for i := 0; i < trials; i++ {
		item, _ := selector.Pick()
		counts[item]++
	}

	// Expected probabilities
	expectedProbs := map[string]float64{
		"Apple":  3.0 / 10.0,
		"Banana": 1.0 / 10.0,
		"Cherry": 6.0 / 10.0,
	}

	// Check if actual probability is within ±5% tolerance
	tolerance := 0.05
	for item, expectedProb := range expectedProbs {
		actualProb := float64(counts[item]) / float64(trials)
		if actualProb < expectedProb-tolerance || actualProb > expectedProb+tolerance {
			t.Errorf("Probability for %s is outside expected range: got %f, expected ~%f", item, actualProb, expectedProb)
		}
	}
}

func BenchmarkNewWeightedSelector(b *testing.B) {
	b.ReportAllocs()
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	for i := 0; i < b.N; i++ {
		_, _ = NewWeightedSelector(weightedItems)
	}
}

func BenchmarkNewAutoWeightedSelector(b *testing.B) {
	b.ReportAllocs()
	items := []string{"Dog", "Cat", "Fish"}

	for i := 0; i < b.N; i++ {
		_, _ = NewAutoWeightedSelector(items)
	}
}

func BenchmarkPick(b *testing.B) {
	b.ReportAllocs()
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	selector, _ := NewWeightedSelector(weightedItems)

	b.ResetTimer() // Ignore setup time in benchmark
	for i := 0; i < b.N; i++ {
		_, _ = selector.Pick()
	}
}
