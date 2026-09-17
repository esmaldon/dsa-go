package main

import (
	"fmt"
	"hash/fnv"
)

type node[K comparable, V any] struct {
	key      K
	value    V
	nextNode *node[K, V]
}

type myHashTable[K comparable, V any] struct {
	array []*node[K, V]
	size  int
}

func newHashTable[K comparable, V any](size int) *myHashTable[K, V] {
	return &myHashTable[K, V]{
		array: make([]*node[K, V], size),
		size:  size,
	}
}

func (m *myHashTable[K, V]) hash(key K) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32)
}

func (m *myHashTable[K, V]) Insert(key K, value V) {

}

func main() {
	var keys = []string{"key1", "key2", "key3"}
	var values = []int{1, 2, 3}

}
