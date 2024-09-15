package hashtable

import "fmt"

const bucketCount = 10

// Node structure for linked list
type Node struct {
	Key   string
	Value int
	Next  *Node
}

// HashTable structure
type HashTable struct {
	buckets []*Node
}

// NewHashTable creates a new hash table
func NewHashTable() *HashTable {
	return &HashTable{
		buckets: make([]*Node, bucketCount),
	}
}

// Hash function
func hash(key string) int {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) % bucketCount
	}

	return hash
}

// Insert adds a key-value pair to the hash table
func (ht *HashTable) Insert(key string, value int) {
	index := hash(key)
	newNode := &Node{Key: key, Value: value}

	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else {
		current := ht.buckets[index]
		for current != nil {
			if current.Key == key {
				current.Value = value
				return
			}
			current = current.Next
		}
		newNode.Next = ht.buckets[index]
		ht.buckets[index] = newNode
	}
}

// Search retrieves a value by key
func (ht *HashTable) Search(key string) (int, bool) {
	index := hash(key)
	current := ht.buckets[index]
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return 0, false
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) bool {
	index := hash(key)
	current := ht.buckets[index]
	var prev *Node

	for current != nil {
		if current.Key == key {
			if prev == nil {
				ht.buckets[index] = current.Next
			} else {
				prev.Next = current.Next
			}
			return true
		}
		prev = current
		current = current.Next
	}
	return false
}

func HashTableExample() {
	fmt.Println("Hash table example")

	ht := NewHashTable()
	ht.Insert("foo", 42)
	ht.Insert("bar", 84)

	if value, found := ht.Search("foo"); found {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}

	ht.Delete("foo")
	if value, found := ht.Search("foo"); found {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}
}