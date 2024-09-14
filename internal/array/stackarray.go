package array

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("unable to pop empty stack")
	}

	topItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return topItem, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("unable to keep empty stack")
	}

	return s.items[len(s.items)-1], nil
}

func StackArrayExample() {
	fmt.Println("Stack array example")
	stack := &Stack{}

	fmt.Println("Size: ", stack.Size())
	fmt.Print("Is it empty? ", stack.IsEmpty())

	stack.Push(1)
	stack.Push(4)

	fmt.Println("Size: ", stack.Size())
	fmt.Println("Is it empty? ", stack.IsEmpty())

	top, _ := stack.Peek()
	fmt.Println("Last item is without removing it: ", top)
	fmt.Println("Size: ", stack.Size())
	popped, _ := stack.Pop()
	fmt.Println("Last item is with removing it: ", popped)
	fmt.Println("Size: ", stack.Size())
}