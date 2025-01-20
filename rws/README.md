# Random Weighted Selection Algorithm in Go

The **Weighted Random Selection** algorithm allows selecting an item from a list where each 
item has an associated weight. The probability of selecting an item is **proportional** to its weight.

This Go package implements an **efficient weighted selection algorithm** optimized using **binary search (O(log n))**.


## 🚀 Features
- **Fast Selection**: Uses **binary search** for **O(log n)** selection time.
- **Generics Support**: Works with any data type (`T`).
- **Two Modes**:
  1. **Explicit Weights**: Provide custom weights.
  2. **Auto Weights**: Assigns random weights automatically.
- **Use Cases**:
  - Load balancing: Distribute requests based on server capacity.
  - Gaming: Loot drop probabilities, AI decision-making.
  - Recommendation Systems: Prioritize items based on scores.
  - Simulations: Randomized event occurrences.

## 📦 Installation
```sh
go get github.com/Ja7ad/algo/rws
```

## 🛠️ Usage

### **1️⃣ Selecting with Custom Weights**
```go
package main

import (
	"fmt"
	"log"
	"github.com/Ja7ad/algo/rws"
)

func main() {
	// Define items with custom weights
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

### **2️⃣ Selecting with Auto-Assigned Weights**
```go
items := []string{"Dog", "Cat", "Fish"}

// Create a selector with auto-assigned random weights
autoSelector, err := rws.NewAutoWeightedSelector(items)
if err != nil {
	log.Fatal(err)
}

// Pick a random item
autoSelectedItem, _ := autoSelector.Pick()
fmt.Println("Auto Selected:", autoSelectedItem)
```

## 📊 Mathematical Formula

### **Given:**
- A set of items: $ S = \{s_1, s_2, ..., s_n\} $
- A corresponding weight for each item: $ W = \{w_1, w_2, ..., w_n\} $
- The total sum of weights: $ W_{\text{sum}} = \sum_{i=1}^{n} w_i $
- A random number $R$ sampled from **$[0, W_{\text{sum}})$**

---

## 🎯 Selection Process

1. Generate a random number $R$ uniformly from **$[0, W_{\text{sum}})$**.
2. Iterate through the items, keeping a cumulative sum: $C_j = \sum_{i=1}^{j} w_i $
3. Select the first item $s_j$ where: $ C_j > R $

---

## 🏎️ **Performance Optimization**
- **Initial Approach (O(n))**:  
  - Iterates through all items to find the selection.
  - **Not ideal for large datasets**.

- **Optimized Approach (O(log n))**:  
  - Uses **binary search** on precomputed cumulative weights.
  - **Faster selection**, making it ideal for large-scale applications.

---

## 🏆 **Example Calculation**
### **Given items and weights:**
$$ \text{Items} = [A, B, C] $$
$$ \text{Weights} = [3, 1, 6] $$

1. Compute total weight:
   $$ W_{\text{sum}} = 3 + 1 + 6 = 10 $$

2. Compute cumulative weights:
   - $C_1 = 3$ (for A)
   - $C_2 = 3 + 1 = 4$ (for B)
   - $C_3 = 3 + 1 + 6 = 10$ (for C)

3. Generate a random number $R \in [0, 10)$:
   - If $0 \leq R < 3$ → Select **A**
   - If $3 \leq R < 4$ → Select **B**
   - If $4 \leq R < 10$ → Select **C**


## 📝 **License**
This project is licensed under the **MIT License**.

## 🤝 **Contributions**
Feel free to **fork** this repository and submit **pull requests**. Any contributions to improve performance or add new selection methods are welcome! 🚀
