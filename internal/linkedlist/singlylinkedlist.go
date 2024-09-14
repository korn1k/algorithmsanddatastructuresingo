package linkedlist

import "fmt"

type Node struct {
	data string
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func (list *SinglyLinkedList) InsertAtEnd(data string) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *SinglyLinkedList) InsertAtBeginning(data string) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	oldHead := list.head
	newNode.next = oldHead
	list.head = newNode
}

func (list *SinglyLinkedList) Delete(data string) error {
	if list.head == nil {
		return fmt.Errorf("unable to delete from empty list")
	}

	if list.head.data == data {
		next := list.head.next
		list.head = next
		return nil
	}

	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next == nil {
		return fmt.Errorf("unable to find item to delete")
	}

	current.next = current.next.next
	return nil
}

func (list *SinglyLinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Printf("%s -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *SinglyLinkedList) DeleteAt(index int) error {
	if list.head == nil {
		return fmt.Errorf("unable to delete from empty list")
	}

	if index == 0 {
		list.head = list.head.next
		return nil
	}

	current := list.head
	for i := 0; current.next != nil && i < index-1; i++ {
		current = current.next
	}

	if current.next == nil {
		return fmt.Errorf("index out of range")
	}

	current.next = current.next.next
	return nil
}

func SinglyLinkedListExample() {
	fmt.Println("Singly linked list example")
	list := &SinglyLinkedList{}

	list.InsertAtBeginning("potato")
	list.InsertAtBeginning("pumpkin")
	list.InsertAtEnd("carrot")
	list.InsertAtEnd("banana")
	list.PrintList()
	err := list.DeleteAt(3)
	if err != nil {
		fmt.Println(err)
	}

	list.PrintList()
	list.Delete("pumpkin")
	list.PrintList()
}