package main

import (
	"fmt"
	"github.com/Ja7ad/algo/rws"
	"log"
)

func main() {
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	selector, err := rws.NewWeightedSelector(weightedItems)
	if err != nil {
		log.Fatal(err)
	}

	selectedItem, _ := selector.Pick()
	fmt.Println("Selected:", selectedItem)

	items := []string{"Dog", "Cat", "Fish"}

	autoSelector, err := rws.NewAutoWeightedSelector(items)
	if err != nil {
		log.Fatal(err)
	}

	autoSelectedItem, _ := autoSelector.Pick()
	fmt.Println("Auto Selected:", autoSelectedItem)
}
