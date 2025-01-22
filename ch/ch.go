package ch

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// Hash function type
type Hash func(data []byte) uint32

// Map represents the consistent hash ring with generics
type Map[T any] struct {
	mu       sync.RWMutex
	hash     Hash
	replicas int
	keys     []int          // Sorted virtual node positions
	hashMap  map[int]string // Virtual node hash -> Real node
	data     map[string]T
}

// New creates a new Consistent Hashing instance
func New[T any](replicas int, fn Hash) *Map[T] {
	m := &Map[T]{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
		data:     make(map[string]T),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// AddNode adds a node to the hash ring
func (m *Map[T]) AddNode(node string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i := 0; i < m.replicas; i++ {
		hash := int(m.hash([]byte(strconv.Itoa(i) + node)))
		m.keys = append(m.keys, hash)
		m.hashMap[hash] = node
	}

	sort.Ints(m.keys)
}

// RemoveNode removes a node from the hash ring
func (m *Map[T]) RemoveNode(node string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var newKeys []int
	for _, hash := range m.keys {
		if m.hashMap[hash] != node {
			newKeys = append(newKeys, hash)
		} else {
			delete(m.hashMap, hash)
		}
	}
	m.keys = newKeys
}

// GetNode returns the closest node for the provided key
func (m *Map[T]) GetNode(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	if idx == len(m.keys) {
		idx = 0
	}
	return m.hashMap[m.keys[idx]]
}

// AddKey stores a key-value pair in the correct node
func (m *Map[T]) AddKey(key string, value T) {
	node := m.GetNode(key)

	// If no node found, no need to store the value
	if node == "" {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// RemoveKey deletes a key from the system
func (m *Map[T]) RemoveKey(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// GetKey retrieves a value stored in the system
func (m *Map[T]) GetKey(key string) (T, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.data[key]
	return value, exists
}
