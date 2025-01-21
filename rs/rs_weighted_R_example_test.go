package rs

import "fmt"

func ExampleWeightedReservoirR() {
	stream := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	weights := []float64{0.1, 0.5, 0.2, 0.3, 0.8, 0.6, 0.9, 0.4, 0.7, 0.2}

	k := 5
	sample := WeightedReservoirR(stream, weights, k)
	fmt.Println(sample)
}
