package binarytree

import "fmt"

type MapNode struct {
    Key   int
    Value string
    Left  *MapNode
    Right *MapNode
}

type BSTMap struct {
    Root *MapNode
}

// Insert adds a key-value pair to the map
func (m *BSTMap) Insert(key int, value string) {
    m.Root = insertNodeMap(m.Root, key, value)
}

// Helper function to insert a node in the BST Map
func insertNodeMap(node *MapNode, key int, value string) *MapNode {
    if node == nil {
        return &MapNode{Key: key, Value: value}
    }
    if key < node.Key {
        node.Left = insertNodeMap(node.Left, key, value)
    } else if key > node.Key {
        node.Right = insertNodeMap(node.Right, key, value)
    } else {
        node.Value = value // Update value if key already exists
    }
    return node
}

// Find retrieves the value for a given key
func (m *BSTMap) Find(key int) (string, bool) {
    node := findNodeMap(m.Root, key)
    if node != nil {
        return node.Value, true
    }
    return "", false
}

// Helper function to search a node in the BST Map
func findNodeMap(node *MapNode, key int) *MapNode {
    if node == nil {
        return nil
    }
    if key == node.Key {
        return node
    }
    if key < node.Key {
        return findNodeMap(node.Left, key)
    }
    return findNodeMap(node.Right, key)
}

// Delete removes a key-value pair from the map
func (m *BSTMap) Delete(key int) {
    m.Root = deleteNodeMap(m.Root, key)
}

// Helper function to delete a node from the BST Map
func deleteNodeMap(node *MapNode, key int) *MapNode {
    if node == nil {
        return nil
    }
    if key < node.Key {
        node.Left = deleteNodeMap(node.Left, key)
    } else if key > node.Key {
        node.Right = deleteNodeMap(node.Right, key)
    } else {
        // Node to be deleted found
        if node.Left == nil {
            return node.Right
        }
        if node.Right == nil {
            return node.Left
        }

        // Node has two children, find the in-order successor
        minRight := findMinMap(node.Right)
        node.Key = minRight.Key
        node.Value = minRight.Value
        node.Right = deleteNodeMap(node.Right, minRight.Key)
    }
    return node
}

// Helper function to find the minimum node in a subtree
func findMinMap(node *MapNode) *MapNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}

func BSTMapExample() {
	fmt.Println("\nBST map example")

    m := &BSTMap{}
    m.Insert(10, "ten")
    m.Insert(5, "five")
    m.Insert(15, "fifteen")
    m.Insert(3, "three")
    m.Insert(7, "seven")

    val, found := m.Find(7)
    fmt.Println("Find 7:", val, found)  // seven true

    val, found = m.Find(4)
    fmt.Println("Find 4:", val, found)  // "" false

    m.Delete(7)
    val, found = m.Find(7)
    fmt.Println("Find 7 after deletion:", val, found)  // "" false
}