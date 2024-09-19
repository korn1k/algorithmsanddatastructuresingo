package binarytree

import "fmt"

type SetNode struct {
    Val   int
    Left  *SetNode
    Right *SetNode
}

type BSTSet struct {
    Root *SetNode
}

// Insert adds a value to the set
func (s *BSTSet) Insert(val int) {
    s.Root = insertNode(s.Root, val)
}

// Helper function to insert a node in the BST
func insertNode(node *SetNode, val int) *SetNode {
    if node == nil {
        return &SetNode{Val: val}
    }
    if val < node.Val {
        node.Left = insertNode(node.Left, val)
    } else if val > node.Val {
        node.Right = insertNode(node.Right, val)
    }
    // If val == node.Val, it already exists, do nothing
    return node
}

// Find checks if a value exists in the set
func (s *BSTSet) Find(val int) bool {
    return findNode(s.Root, val)
}

// Helper function to search a node in the BST
func findNode(node *SetNode, val int) bool {
    if node == nil {
        return false
    }
    if val == node.Val {
        return true
    }
    if val < node.Val {
        return findNode(node.Left, val)
    }
    return findNode(node.Right, val)
}

// Delete removes a value from the set
func (s *BSTSet) Delete(val int) {
    s.Root = deleteNode(s.Root, val)
}

// Helper function to delete a node from the BST
func deleteNode(node *SetNode, val int) *SetNode {
    if node == nil {
        return nil
    }
    if val < node.Val {
        node.Left = deleteNode(node.Left, val)
    } else if val > node.Val {
        node.Right = deleteNode(node.Right, val)
    } else {
        // Node to be deleted found
        if node.Left == nil {
            return node.Right
        }
        if node.Right == nil {
            return node.Left
        }

        // Node has two children, find the in-order successor
        minRight := findMinn(node.Right)
        node.Val = minRight.Val
        node.Right = deleteNode(node.Right, minRight.Val)
    }
    return node
}

// Helper function to find the minimum node in a subtree
func findMinn(node *SetNode) *SetNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}

func BSTSetExample() {
    set := &BSTSet{}
    set.Insert(10)
    set.Insert(5)
    set.Insert(15)
    set.Insert(3)
    set.Insert(7)

    fmt.Println("Find 7:", set.Find(7))   // true
    fmt.Println("Find 4:", set.Find(4))   // false

    set.Delete(7)
    fmt.Println("Find 7 after deletion:", set.Find(7))  // false
}