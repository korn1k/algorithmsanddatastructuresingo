package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/korn1k/algorithmsanddatastructuresingo/internal/array"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/backtracking"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/binarysearch"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/binarytree"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/hashtable"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/linkedlist"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/recursion"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/sort"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Arrays")
        fmt.Println("2. Stacks")
        fmt.Println("3. Sort algorithms")
        fmt.Println("4. Recursion")
        fmt.Println("5. Hashtable")
        fmt.Println("6. Binary search")
        fmt.Println("7. Binary tree")
        fmt.Println("8. Backtracking")
        fmt.Println("9. Exit")
        fmt.Print("Enter your choice by specifying number: ")

        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        choice, err := strconv.Atoi(input[:len(input)-1])
        if err != nil {
            fmt.Println("Invalid choice. Please enter a valid number.")
            continue
        }

        switch choice {
        case 1:
            array.StaticArrayExample()
			array.DynamicArrayExample()
			array.StackArrayExample()
        case 2:
            linkedlist.SinglyLinkedListExample()
            linkedlist.DoublyLinkedListExample()
            linkedlist.QueueLinkedListExample()
        case 3:
            sort.InsertSortExample()
            sort.MergeSortExample()
            sort.QuickSortExample()
            sort.BucketSortExample()
        case 4:
            recursion.FactorialRecursionExample()
            recursion.FibonacciRecursionExample()
        case 5:
            hashtable.HashTableExample()
        case 6:
            binarysearch.ArrayBinarySearchExample()
        case 7:
            binarytree.BinaryTreeExample()
            binarytree.BSTExample()
            binarytree.BFSExample()
            binarytree.BSTMapExample()
            binarytree.BSTSetExample()
        case 8:
            backtracking.BinaryTreeDFSExample()
        case 9:
            fmt.Println("Exiting from the selection")
            return
        default:
            panic("not implemented")
        }
    }
}