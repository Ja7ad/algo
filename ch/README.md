# Consistent Hashing

In computer science, consistent hashing is a special kind of hashing technique such that when a hash table is resized, only 
$\displaystyle n/m$ keys need to be remapped on average where $\displaystyle n$ is the number of keys and
$\displaystyle m$ is the number of slots. In contrast, in most traditional hash tables, a change in the number of array 
slots causes nearly all keys to be remapped because the mapping between the keys and the slots is defined by a modular operation.

## Project used algorithm

- Couchbase automated data partitioning
- OpenStack's Object Storage Service Swift
- Partitioning component of Amazon's storage system Dynamo
- Data partitioning in Apache Cassandra
- Data partitioning in ScyllaDB
- Data partitioning in Voldemort
- Akka's consistent hashing router
- Riak, a distributed key-value database
- Gluster, a network-attached storage file system
- Akamai content delivery network
- Discord chat application
- Load balancing gRPC requests to a distributed cache in SpiceDB
- Chord algorithm
MinIO object storage system

## 📊 **Mathematical Formula for Consistent Hashing**

### **Problem Definition**
Given a set of `N` nodes and `K` keys, we need to distribute the keys among the nodes **such that minimal data movement is required** when nodes are added or removed.

### **Hash Ring Representation**
1. We define a **circular space** from `0` to `M-1`, where `M = 2^m` for an `m`-bit hash function.
2. Each **node** `n_i` is hashed using function `H(n_i)`, assigning it a position on the ring: $P(n_i) = H(n_i) \mod M$
3. Each **key** `k_j` is hashed to the ring using the same function: $P(k_j) = H(k_j) \mod M$
4. A **key is assigned to the first node encountered in the clockwise direction** from its position.

### **Mathematical Proof of Load Balancing**
The expected number of keys per node is given by: $E[\text{keys per node}] = \frac{K}{N}$
where:
- `K` is the total number of keys.
- `N` is the total number of nodes.

If a node **joins**, it takes responsibility for keys previously mapped to the **next node**, meaning only: $\frac{K}{N+1}$
keys are affected, significantly reducing data movement compared to traditional hashing (`O(K)` movement).
If a node **leaves**, its keys are reassigned to the **next available node**, again affecting only: $\frac{K}{N-1}$
keys instead of `O(K)`.

### **Time Complexity**
| Operation         | Complexity |
|-------------------|------------|
| **Node Addition** | `O(K/N + log N)` |
| **Node Removal**  | `O(K/N + log N)` |
| **Key Lookup**    | `O(log N)` (Binary Search) |
| **Add a key**     | `O(log N)`|
| **Remove a key**    | `O(log N)` |



## 🧪 **Mathematical Test Case for Consistent Hashing**
### **Test Case Design**
To validate **Consistent Hashing**, we check:
1. **Keys are evenly distributed** across nodes (`K/N` per node).
2. **Minimal keys move on node addition/removal** (`K/N+1` or `K/N-1`).
3. **Lookups are efficient (`O(log N)`)** using binary search.

### **Example**
#### **Initial Nodes (`N = 3`)**
| Node | Hash Value (Position on Ring) |
|------|-----------------------------|
| `A`  | `H(A) = 15` |
| `B`  | `H(B) = 45` |
| `C`  | `H(C) = 90` |

#### **Keys (`K = 6`)**
| Key  | Hash Value | Assigned Node |
|------|-----------|--------------|
| `k1` | `H(k1) = 10` | `A` |
| `k2` | `H(k2) = 30` | `B` |
| `k3` | `H(k3) = 55` | `C` |
| `k4` | `H(k4) = 70` | `C` |
| `k5` | `H(k5) = 85` | `C` |
| `k6` | `H(k6) = 95` | `A` |

#### **After Adding `Node D (H(D) = 60)`**
Only **`k3` and `k4`** move to `D`, while other keys remain unaffected.
