package rs

import "fmt"

func ExampleReservoirSampleR() {
	// Define a sample stream of integers
	stream := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Select 5 random elements using Algorithm R
	reservoir := ReservoirSampleR(stream, 5)

	// Print the selected reservoir sample
	fmt.Println("Selected Reservoir Sample:", reservoir)
}
