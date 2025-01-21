package rws

import (
	"fmt"
	"log"
)

func ExampleNewWeightedSelector() {
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	selector, err := NewWeightedSelector(weightedItems)
	if err != nil {
		log.Fatal(err)
	}

	selectedItem, _ := selector.Pick()
	fmt.Println("Selected:", selectedItem)
}

func ExampleNewAutoWeightedSelector() {
	items := []string{"Dog", "Cat", "Fish"}

	autoSelector, err := NewAutoWeightedSelector(items)
	if err != nil {
		log.Fatal(err)
	}

	autoSelectedItem, _ := autoSelector.Pick()
	fmt.Println("Auto Selected:", autoSelectedItem)
}
