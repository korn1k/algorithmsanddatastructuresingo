package binarytree

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val, Left: nil, Right: nil}
}

// Root -> Left -> Right
func preOrder(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("%d ", root.Val)
    preOrder(root.Left)
    preOrder(root.Right)
}

// Left -> Root -> Right
func inOrder(root *TreeNode) {
    if root == nil {
        return
    }
    inOrder(root.Left)
    fmt.Printf("%d ", root.Val)
    inOrder(root.Right)
}

// Left -> Right -> Root
func postOrder(root *TreeNode) {
    if root == nil {
        return
    }
    postOrder(root.Left)
    postOrder(root.Right)
    fmt.Printf("%d ", root.Val)
}

func BinaryTreeExample() {
	fmt.Println("Binary tree example")

    // Creating the following binary tree:
    //       1
    //      / \
    //     2   3
    //    / \   \
    //   4   5   6

    root := newNode(1)
    root.Left = newNode(2)
    root.Right = newNode(3)
    root.Left.Left = newNode(4)
    root.Left.Right = newNode(5)
    root.Right.Right = newNode(6)

    fmt.Println("Preorder Traversal:")
    preOrder(root)  // Output: 1 2 4 5 3 6

    fmt.Println("\nInorder Traversal:")
    inOrder(root)   // Output: 4 2 5 1 3 6

    fmt.Println("\nPostorder Traversal:")
    postOrder(root) // Output: 4 5 2 6 3 1
}