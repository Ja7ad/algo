# Reservoir sampling

Reservoir sampling is a family of randomized algorithms for choosing a simple random sample, without replacement, 
of k items from a population of unknown size n in a single pass over the items. The size of the population n is not 
known to the algorithm and is typically too large for all n items to fit into main memory. The population is revealed 
to the algorithm over time, and the algorithm cannot look back at previous items. At any point, the current state of 
the algorithm must permit extraction of a simple random sample without replacement of size k over the part of 
the population seen so far.

## **🔹 Variants of Reservoir Sampling**
While **Algorithm R** is the simplest and most commonly used, there are **other variants** that improve performance in specific cases:

| **Algorithm**                 | **Description** | **Complexity** |
|--------------------------------|----------------|---------------|
| **Algorithm R**               | Basic reservoir sampling, replaces elements with probability `k/i` | **O(N)** |
| **Algorithm L**               | Optimized for large `N`, reduces replacements via skipping | **O(N), fewer iterations** |
| **Weighted Reservoir Sampling** | Assigns elements weights, prioritizing selection based on weight | **O(N log k)** (heap-based) |
| **Random Sort Reservoir Sampling** | Uses a min-heap priority queue, selecting `k` elements with highest random priority scores | **O(N log k)** |

## Algorithm Weighted R – Weighted Reservoir Sampling
**Weighted Reservoir Sampling** is an **efficient algorithm** for selecting `k` elements **proportionally to their weights** from a stream of unknown length `N`, using only `O(k)` memory.  

This repository implements **Weighted Algorithm R**, an extension of **Jeffrey Vitter's Algorithm R**, which allows weighted sampling using a **heap-based approach**.

> This algorithm uses a **min-heap-based priority selection**, ensuring **O(N log k)** time complexity, making it efficient for large streaming datasets.

## 📊 **Mathematical Formula for Weighted Algorithm R**

### **Problem Definition**
We need to select **`k` elements** from a data stream **of unknown length `N`**, ensuring **each element is selected with a probability proportional to its weight `w_i`**.

### **Algorithm Steps**
1. **Initialize a Min-Heap of Size `k`**
   - Store the first `k` elements **with their priority scores**:
     \[
     $p_i = \frac{w_i}{U_i}$
     \]
     where \( $U_i$ \) is a uniform random number from **(0,1]**.

2. **Process Remaining Elements (`i > k`)**
   - For each new element `s_i`:
     - Compute **priority score**:
       \[
       $p_i = \frac{w_i}{U_i}$
       \]
     - If `p_i` is greater than the **smallest priority in the heap**, replace the smallest element.

3. **After processing `N` elements**, the reservoir will contain `k` elements **selected proportionally to their weights**.

---

## 🔬 **Probability Proof**
For any element \( $s_i$ \) with weight \( $w_i$ \):
1. The **priority score** is:
   \[
   $p_i = \frac{w_i}{U_i}$
   \]
   where \( $U_i \sim U(0,1]$ \).

2. The **probability that `s_i` is among the top `k` elements**:
   \[
   $P(s_i \text{ is selected}) \propto w_i$
   \]
   meaning elements with **higher weights** are **more likely to be selected**.

✅ **Conclusion:** Weighted Algorithm R correctly samples elements **proportionally to their weights**, unlike uniform Algorithm R.

---

## 🧪 **Test Case Formula for Weighted Algorithm R**

### **Test Case Design**
To validate Weighted Algorithm R, we must check if:
- **Higher-weight elements are chosen more frequently**.
- **Selection follows the weight distribution over multiple runs**.

### **Mathematical Test**
For `T` independent runs:
- Let `count(s_i)` be the number of times `s_i` appears in the reservoir.
- Expected probability:
  \[
  $P(s_i) = \frac{w_i}{\sum w_j}$
  \]
- Expected occurrence over `T` runs:
  \[
  $\text{Expected count}(s_i) = T \times \frac{w_i}{\sum w_j}$
  \]
- We verify that `count(s_i)` is **statistically close** to this value.

# 🎯 Algorithm L

**Reservoir Sampling** is a technique for randomly selecting `k` elements from a stream of unknown length `N`.  
**Algorithm L**, introduced by **Jeffrey Vitter (1985)**, improves upon traditional methods by using an **optimized skipping approach**, significantly reducing the number of random number calls.

### **Problem Definition**
We need to select **`k` elements** from a data stream **of unknown length `N`**, ensuring **each element has an equal probability `k/N`** of being chosen.

### **Algorithm Steps**
1. **Fill the reservoir** with the **first `k` elements**.  
2. **Initialize weight factor `W`** using:

   $W = \exp\left(\frac{\log(\text{random}())}{k}\right)$

3. **Skip elements efficiently** using the geometric formula:
   
   $\text{skip} = \lfloor \frac{\log(\text{random}())}{\log(1 - W)} \rfloor$

4. **If still in bounds**, **randomly replace** an element in the reservoir.  
5. **Update `W`** for the next iteration using:

   $W = W \times \exp\left(\frac{\log(\text{random}())}{k}\right)$

6. **Repeat until the end of the stream**.

### **Probability Proof**
For each element \( $s_i$ \), we show that it has an equal probability of being selected:

1. The probability that \( $s_i$ \) **reaches the selection process**:

   $P(s_i \text{ is considered}) = \frac{k}{i}$

2. The probability that \( $s_i$ \) **remains in the reservoir** is:

   $P(s_i \text{ in final reservoir}) = \frac{k}{N}, \quad \forall i \in \{1, ..., N\}$

This confirms that **Algorithm L ensures uniform selection**.


## 🧪 **Test Case Formula for Algorithm L**

### **Test Case Design**
To validate Algorithm L, we must check if:
- **Each element is chosen with probability `k/N`**.
- **Selection is uniform over multiple runs**.

### **Mathematical Test**
For `T` independent runs:
- Let `count(s_i)` be the number of times `s_i` appears in the reservoir.
- Expected probability:
  
   $P(s_i) = \frac{k}{N}$

- Expected occurrence over `T` runs:

  $\text{Expected count}(s_i) = T \times \frac{k}{N}$

- We verify that `count(s_i)` is **statistically close** to this value.
