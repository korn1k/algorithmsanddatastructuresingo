package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/korn1k/algorithmsanddatastructuresingo/internal/array"
	"github.com/korn1k/algorithmsanddatastructuresingo/internal/linkedlist"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Arrays")
        fmt.Println("2. Stacks")
        fmt.Println("3. Exit")
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
        case 3:
            fmt.Println("Exiting from the selection")
            return
        default:
            panic("not implemented")
        }
    }
}