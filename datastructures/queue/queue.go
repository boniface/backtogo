package main

import "fmt"

//Create Node
type Node struct {
	value int
	next  *Node
}

//Create a Queue
type Queue struct {
	head *Node
	tail *Node
	size int
}

//Get Size of the Queue
func (queue *Queue) Size() int {
	return queue.size
}

// Check if Queue is Empty
func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

// Get Head Value for Queue
func (queue *Queue) Peek() (int, bool) {
	if queue.IsEmpty() {
		fmt.Println("Queue ")
		return 0, false
	}
	return queue.head.value, true
}

// Print the Contents of the Queue
func (queue *Queue) Print() {
	temp := queue.head
	for temp != nil {
		fmt.Println("Queue Value ", temp.value)
		temp = temp.next
	}
}

func (queue *Queue) Add(value int) {
	// Create a New Node to Add
	newNode := &Node{value, nil}
	// Check if Node has a NO head
	if queue.head == nil {
		// Then Add New Node as Head
		queue.head = newNode

		// Make it As tail As well
		queue.tail = newNode
	} else {
		// Other Wise Que has Nodes in It
		// Add newNode to the Tail of the Current Tail
		queue.tail.next = newNode
		//Then Make the Tail the New Node
		queue.tail = newNode
	}
	// Increase the Size of the queue by 1
	queue.size++
}

func (queue *Queue) Remove() (int, bool) {
	// Check if Queue is empty
	if queue.IsEmpty() {
		fmt.Println("Queue is Empty")
		return 0, false
	}
	// You only remove from the head in a queue
	// Remove the Value of the head
	removedValue := queue.head.value
	// Make the Queue Head Point to the Next Head
	queue.head = queue.head.next
	// Decrease or Shrink the size of the Queue by 1
	queue.size--
	return removedValue, true
}

func main() {
	// Initilise Some Nodes
	queue := new(Queue)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Print()

	for queue.IsEmpty() == false {
		val, res := queue.Remove()
		fmt.Println("The Value Removed ", val, " Removed? ", res, " Size of Queue", queue.size)
	}
}
