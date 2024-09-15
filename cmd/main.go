package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/korn1k/algorithmsanddatastructuresingo/internal/array"
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
        fmt.Println("5. Exit")
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
            fmt.Println("Exiting from the selection")
            return
        default:
            panic("not implemented")
        }
    }
}