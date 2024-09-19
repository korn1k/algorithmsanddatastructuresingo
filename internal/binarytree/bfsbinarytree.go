package binarytree

import (
    "fmt"
)

type Node struct {
    Val   int
    Left  *Node
    Right *Node
}

// Breadth-First Search (Level-order traversal)
func bfs(root *Node) {
    if root == nil {
        return
    }

    // Create a queue and enqueue the root
    queue := []*Node{root}

    for len(queue) > 0 {
        // Dequeue the front node
        current := queue[0]
        queue = queue[1:]

        // Process the current node (print its value)
        fmt.Println(current.Val)

        // Enqueue the left child if it exists
        if current.Left != nil {
            queue = append(queue, current.Left)
        }

        // Enqueue the right child if it exists
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
}

func BFSExample() {
	fmt.Println("\nBFS example")
    // Create a sample binary tree
    /*
            1
           / \
          2   3
         / \   \
        4   5   6
    */
    root := &Node{Val: 1}
    root.Left = &Node{Val: 2}
    root.Right = &Node{Val: 3}
    root.Left.Left = &Node{Val: 4}
    root.Left.Right = &Node{Val: 5}
    root.Right.Right = &Node{Val: 6}

    // Perform BFS
    fmt.Println("Breadth-First Search (BFS) traversal:")
    bfs(root)
}