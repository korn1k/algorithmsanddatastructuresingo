package binarytree

/*

1. Full binary tree:
        1
       / \
      2   3
     / \  
    4   5 
- Every node has either 0 or 2 children.

2. Complete Binary Tree:
        1
       / \
      2   3
     / \  /
    4   5 6
- All levels are completely filled except possibly the last, and all nodes are as far left as possible.

3. Perfect Binary Tree:
        1
       / \
      2   3
     / \ / \
    4  5 6  7
- All internal nodes have two children, and all leaf nodes are at the same level.

4. Balanced Binary Tree:
        1
       / \
      2   3
     /     \
    4       5
- The difference in heights between the left and right subtrees of any node is at most 1.

5. Binary Search Tree (BST):
        5
       / \
      3   7
     / \ / \
    2  4 6  8
- For every node, the values in the left subtree are less than the node's value, and the values in the right subtree are greater.

*/