# algo: A Collection of High-Performance Algorithms

[![Go Reference](https://pkg.go.dev/badge/github.com/Ja7ad/algo.svg)](https://pkg.go.dev/github.com/Ja7ad/algo)
[![codecov](https://codecov.io/gh/Ja7ad/algo/graph/badge.svg?token=9fLKrkUviU)](https://codecov.io/gh/Ja7ad/algo)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ja7ad/algo)](https://goreportcard.com/report/github.com/Ja7ad/algo)

**`algo`** is a Golang library featuring a variety of **efficient** and **well-optimized** algorithms designed for diverse **problem-solving needs**. 

## 📌 Features
✅ **Optimized Performance** – Algorithms are designed with efficiency in mind.  
✅ **Modular Structure** – Each algorithm is in its own package for easy use.  
✅ **Well-Documented** – Clear documentation and examples for every algorithm.  
✅ **Tested & Benchmarked** – Includes comprehensive tests and benchmarks.  


## 📚 Available Algorithms

| Algorithm | Description |
|-----------|-------------|
| [Random Weighted Selection](./rws/README.md) | Selects items randomly based on assigned weights. Useful in load balancing, gaming, and AI. |


## 🚀 Installation >= go 1.19

```shell
go get -u github.com/Ja7ad/algo
```

## ✅ Usage Example

Here’s how you can use the **Random Weighted Selection** algorithm:

```go
package main

import (
	"fmt"
	"log"
	"github.com/Ja7ad/algo/rws"
)

func main() {
	// Define items with weights
	weightedItems := map[int]string{
		3: "Apple",
		1: "Banana",
		6: "Cherry",
	}

	// Create a selector
	selector, err := rws.NewWeightedSelector(weightedItems)
	if err != nil {
		log.Fatal(err)
	}

	// Pick a random item
	selectedItem, _ := selector.Pick()
	fmt.Println("Selected:", selectedItem)
}
```

For more details, check the [Random Weighted Selection documentation](./rws/README.md).

## 📌 Contribution

We welcome contributions! Feel free to submit pull requests, open issues, or suggest new algorithms.

## 📜 License

This project is licensed under the **MIT License**.