package linkedlist

import "fmt"

type DoublyNode struct {
	data string
	next *DoublyNode
	previous *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
}

func (list *DoublyLinkedList) InsertAtBeginning(data string) {
	newNode := &DoublyNode{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	oldHead := list.head
	oldHead.previous = newNode
	newNode.next = oldHead
	list.head = newNode
}

func (list *DoublyLinkedList) InsertAtEnd(data string) {
	newNode := &DoublyNode{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	oldTail := list.tail
	newNode.previous = oldTail
	oldTail.next = newNode
	list.tail = newNode
}

func (list *DoublyLinkedList) Delete(data string) error {
	if list.head == nil {
		return fmt.Errorf("unable to delete because the list is empty")
	}

	if list.head.data == data {
		next := list.head.next

		if next == nil {
			list.head = nil
			list.tail = nil
		} else {
			next.previous = nil
			list.head = next
		}
	}

	if list.tail.data == data {
		previous := list.tail.previous

		if previous == nil {
			list.head = nil
			list.tail = nil
		} else {
			previous.next = nil
			list.tail = previous
		}
	}

	current := list.head
	for current != nil && current.data != data {
		current = current.next
	}

	if current == nil {
		return fmt.Errorf("unable to delete non-existing")
	}

	if current.previous != nil {
		current.previous.next = current.next
	}

	if current.next != nil {
		current.next.previous = current.previous
	} 

	return nil
}

func (list *DoublyLinkedList) DisplayListBackward() {
	current := list.tail
	for current != nil {
		fmt.Printf("%s <-> ", current.data)
		current = current.previous
	}
	fmt.Println("nil")
}

func DoublyLinkedListExample() {
	fmt.Println("Doubly linked list example")
	list := &DoublyLinkedList{}

	list.InsertAtBeginning("yellow")
	list.InsertAtBeginning("black")
	list.InsertAtEnd("purple")
	list.InsertAtEnd("orange")
	list.DisplayListBackward()
	list.Delete("orange")
	list.DisplayListBackward()
	list.Delete("yellow")
	list.DisplayListBackward()
}