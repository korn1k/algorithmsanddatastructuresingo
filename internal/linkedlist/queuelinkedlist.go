package linkedlist

import "fmt"

type QueueNode struct {
	data string
	next *QueueNode
}

type QueueLinkedList struct {
	front *QueueNode
	rear *QueueNode
	size int
}

func (queue *QueueLinkedList) Enqueue(data string) {
	newNode := &QueueNode{data: data}

	if queue.rear == nil {
		queue.rear = newNode
		queue.front = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}

	queue.size++
}

func (queue *QueueLinkedList) Dequeue() (string, error) {
	if queue.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}

	data := queue.front.data
	queue.front = queue.front.next

	if queue.front == nil {
		queue.rear = nil
	}

	queue.size--

	return data, nil
}

func (queue *QueueLinkedList) Peek() (string, error) {
	if queue.IsEmpty() {
		return "", fmt.Errorf("queue is empty")
	}
	return queue.front.data, nil
}

func (queue *QueueLinkedList) IsEmpty() bool {
	return queue.size == 0
}

func (queue *QueueLinkedList) Size() int {
	return queue.size
}

func QueueLinkedListExample() {
	fmt.Println("Queue linked list example")

	queue := &QueueLinkedList{}
	queue.Enqueue("sam")
	queue.Enqueue("patrick")
	fmt.Println("Size: ", queue.Size())
	peek, _ := queue.Peek()
	fmt.Println("Peek: ", peek)
	queue.Dequeue()
	fmt.Println("Size: ", queue.Size())
	peek, _ = queue.Peek()
	fmt.Println("Peek: ", peek)
	queue.Dequeue()
	fmt.Println("Is empty? ", queue.IsEmpty())
}