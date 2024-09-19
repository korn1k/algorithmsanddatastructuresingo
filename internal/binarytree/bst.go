package binarytree

import "fmt"

// For definition of a TreeNode refer to binarytree.go

// Insert a value into the BST
func insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insert(root.Left, val)
    } else {
        root.Right = insert(root.Right, val)
    }
    return root
}

// Search for a value in the BST
func search(root *TreeNode, val int) bool {
    if root == nil {
        return false
    }
    if val == root.Val {
        return true
    }
    if val < root.Val {
        return search(root.Left, val)
    }
    return search(root.Right, val)
}

func findMin(root *TreeNode) *TreeNode {
    current := root
    for current.Left != nil {
        current = current.Left
    }
    return current
}

// Remove a node from the BST
func remove(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    
    // Traverse the tree to find the node to be removed
    if val < root.Val {
        root.Left = remove(root.Left, val)
    } else if val > root.Val {
        root.Right = remove(root.Right, val)
    } else {
        // Node to be removed found

        // Case 1: Node has no children
        if root.Left == nil && root.Right == nil {
            return nil
        }

        // Case 2: Node has one child
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }

        // Case 3: Node has two children
        // Find the in-order successor (minimum value in the right subtree)
        minRight := findMin(root.Right)
        root.Val = minRight.Val
        root.Right = remove(root.Right, minRight.Val)  // Remove the in-order successor
    }
    
    return root
}

func BSTExample() {
	fmt.Println("\nBST example")

    var root *TreeNode
    values := []int{15, 10, 20, 8, 12, 25}
    
    // Insert values into the BST
    for _, val := range values {
        root = insert(root, val)
    }

    fmt.Println("Inorder Traversal:")
    inOrder(root)  // Output: 8 10 12 15 20 25

    // Search for a value
    fmt.Println("\nSearching for 12:", search(root, 12))  // Output: true
    fmt.Println("Searching for 5:", search(root, 5))      // Output: fals
	
	// Remove a node from the BST
    root = remove(root, 20)

    fmt.Println("\nInorder Traversal after removal:")
    inOrder(root)  // Output: 8 10 12 15 25
}
