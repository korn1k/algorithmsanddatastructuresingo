package backtracking

import (
	"fmt"
)

// Define a structure for the tree nodes
type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

// Function to perform backtracking in a tree to find a path to a target node
func backtrack(root *Node, target string, path *[]string) bool {
	// Base case: if the node is nil, return false (no path)
	if root == nil {
		return false
	}

	// Add the current node to the path
	*path = append(*path, root.Val)

	// If the current node is the target, return true
	if root.Val == target {
		return true
	}

	// Recursively try left and right subtrees
	if backtrack(root.Left, target, path) || backtrack(root.Right, target, path) {
		return true
	}

	// If neither subtree leads to the target, backtrack by removing the last node
	*path = (*path)[:len(*path)-1]
	return false
}

func BinaryTreeDFSExample() {
	fmt.Println("\nBinary Tree DFS example")
	// Create a binary tree
	/*
		           A
		         /   \
		        B     C
		       / \   / \
		      D   E F   G
	*/
	root := &Node{Val: "A"}
	root.Left = &Node{Val: "B"}
	root.Right = &Node{Val: "C"}
	root.Left.Left = &Node{Val: "D"}
	root.Left.Right = &Node{Val: "E"}
	root.Right.Left = &Node{Val: "F"}
	root.Right.Right = &Node{Val: "G"}

	// Backtracking to find the path to target node 'E'
	target := "E"
	path := []string{}
	found := backtrack(root, target, &path)

	if found {
		fmt.Printf("Path to node '%s': %v\n", target, path)
	} else {
		fmt.Printf("Node '%s' not found in the tree.\n", target)
	}
}
