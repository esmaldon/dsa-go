package structures

import (
	"fmt"
	"hash/fnv"
)

type hashNode[K comparable, V any] struct {
	key      K
	value    V
	nextNode *hashNode[K, V]
}

type myHashTable[K comparable, V any] struct {
	array     []*hashNode[K, V]
	size      int
	itemCount int
}

func NewHashTable[K comparable, V any](size int) *myHashTable[K, V] {
	return &myHashTable[K, V]{
		array: make([]*hashNode[K, V], size),
		size:  size,
	}
}

func (m *myHashTable[K, V]) hash(key K) int {
	// convert key to string to hash it
	keyStr := fmt.Sprintf("%v", key)

	h := fnv.New32a()
	h.Write([]byte(keyStr))
	// return the residuo to fix the array size
	return int(h.Sum32()) % m.size
}

func (m *myHashTable[K, V]) Insert(key K, value V) {
	index := m.hash(key)
	fmt.Printf("hash code index %v\n", index)
	node := m.array[index]

	for node != nil {
		// if key is already present then only update the value
		if node.key == key {
			node.value = value
			return
		}
		node = node.nextNode
	}
	// create a new hash node and set it in the head of the index array
	newNode := &hashNode[K, V]{
		key:      key,
		value:    value,
		nextNode: m.array[index],
	}
	m.array[index] = newNode
	m.itemCount++
}

func (m *myHashTable[K, V]) Search(key K) (bool, V) {
	index := m.hash(key)
	node := m.array[index]

	for node != nil {
		if node.key == key {
			return true, node.value
		}
		node = node.nextNode
	}
	var zero V
	return false, zero
}

func (m *myHashTable[K, V]) Delete(key K) bool {
	index := m.hash(key)
	node := m.array[index]
	var previous *hashNode[K, V]

	for node != nil {
		if node.key == key {
			if previous == nil && node.nextNode == nil {
				// only head node is in the table index
				m.array[index] = nil
				return true
			}
			if previous == nil && node.nextNode != nil {
				// head is the node to delete but it has a next node
				m.array[index] = node.nextNode
				return true
			}
			previous.nextNode = node.nextNode
			return true
		}
		previous = node
		node = node.nextNode
	}
	return false
}

func (m *myHashTable[K, V]) ItemCount() int {
	return m.itemCount
}

func (m *myHashTable[K, V]) PrintHashTable() {
	fmt.Printf("item count %v\n", m.itemCount)
	for i, node := range m.array {
		fmt.Printf("At index %v of the table, the values are:\n", i)
		for node != nil {
			fmt.Printf("%v\n", node.value)
			node = node.nextNode
		}
	}
}
