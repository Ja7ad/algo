package ch

import (
	"strconv"
	"testing"
)

func TestConsistentHashing_NodeAddition(t *testing.T) {
	ch := New[string](3, nil)
	ch.AddNode("NodeA")
	ch.AddNode("NodeB")
	ch.AddNode("NodeC")

	key := "my-key"
	node := ch.GetNode(key)

	if node == "" {
		t.Errorf("Expected a valid node, but got an empty string")
	}

	sameNode := ch.GetNode(key)
	if node != sameNode {
		t.Errorf("Expected consistent mapping, but got different results")
	}
}

func TestConsistentHashing_AddGetKey(t *testing.T) {
	ch := New[int](3, nil)
	ch.AddNode("NodeA")
	ch.AddNode("NodeB")

	ch.AddKey("user123", 99)
	ch.AddKey("user456", 42)

	value, exists := ch.GetKey("user123")
	if !exists || value != 99 {
		t.Errorf("Expected 99, but got %d", value)
	}

	value, exists = ch.GetKey("user456")
	if !exists || value != 42 {
		t.Errorf("Expected 42, but got %d", value)
	}
}

func TestConsistentHashing_RemoveKey(t *testing.T) {
	ch := New[string](3, nil)
	ch.AddNode("NodeA")
	ch.AddKey("user123", "Data1")

	ch.RemoveKey("user123")

	_, exists := ch.GetKey("user123")
	if exists {
		t.Errorf("Expected key to be removed, but it still exists")
	}
}

type TestStruct struct {
	Name  string
	Score int
}

func TestConsistentHashing_WithStruct(t *testing.T) {
	ch := New[TestStruct](3, nil)
	ch.AddNode("NodeA")
	ch.AddNode("NodeB")

	data := TestStruct{Name: "Alice", Score: 100}
	ch.AddKey("user123", data)

	retrieved, exists := ch.GetKey("user123")
	if !exists || retrieved.Name != "Alice" || retrieved.Score != 100 {
		t.Errorf("Expected Alice with score 100, but got %+v", retrieved)
	}
}

func BenchmarkConsistentHashing_AddNode(b *testing.B) {
	ch := New[string](100, nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.AddNode("Node" + strconv.Itoa(i))
	}
}

func BenchmarkConsistentHashing_RemoveNode(b *testing.B) {
	ch := New[string](100, nil)
	for i := 0; i < 1000; i++ {
		ch.AddNode("Node" + strconv.Itoa(i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.RemoveNode("Node" + strconv.Itoa(i%1000))
	}
}

func BenchmarkConsistentHashing_GetNode(b *testing.B) {
	ch := New[string](100, nil)
	for i := 0; i < 1000; i++ {
		ch.AddNode("Node" + strconv.Itoa(i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ch.GetNode("key" + strconv.Itoa(i))
	}
}

func BenchmarkConsistentHashing_AddKey(b *testing.B) {
	ch := New[string](100, nil)
	for i := 0; i < 1000; i++ {
		ch.AddNode("Node" + strconv.Itoa(i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.AddKey("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}
}

func BenchmarkConsistentHashing_RemoveKey(b *testing.B) {
	ch := New[string](100, nil)
	for i := 0; i < 1000; i++ {
		ch.AddNode("Node" + strconv.Itoa(i))
	}
	for i := 0; i < 10000; i++ {
		ch.AddKey("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.RemoveKey("key" + strconv.Itoa(i%10000))
	}
}
