package main

import "fmt"

// Node Structure
type Node struct {
	IPAddress   string // e.g "10.10.10.9"
	ClassName   string // e.g "Class A"
	LastIPOctet int
	next        *Node
}

// The Stuck Structure
type Stack struct {
	head *Node
	size int
}

// ADTS for the Stack Size(), Push(), Pop(), Print(), IsEmpty(),

// Get the Size of the Stack
func (stack *Stack) Size() int {
	return stack.size
}

// Is the Stack Empty

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) Peek() (*Node, bool) {
	if stack.IsEmpty() {
		fmt.Println(" the Stack is Empty")
		return nil, false
	}
	return stack.head, true
}

func (stack *Stack) Push(value *Node) {
	var node = &Node{value.IPAddress,
		value.ClassName,
		value.LastIPOctet,
		stack.head}
	stack.head = node
	stack.size++
}

func (stack *Stack) Pop() (Node, bool) {
	//Check if Stack is Empty
	if stack.IsEmpty() {
		fmt.Println("Stack is Empty ")
	}
	//Get the Stack value
	value := stack.head
	//Update the Head to the next Top Remaining Stack
	stack.head = stack.head.next
	//Reduce the Size of the Stack
	stack.size--

	//Return the Results
	return *value, true
}

func (stack *Stack) Print() {
	// Put the Stack Head into the temporary Node
	temp := stack.head
	// Print Value
	fmt.Println("Octet Value put in the Stack ::")
	// If the Temp had NO Error
	for temp != nil {
		// Print the Value of the Node Stored in Temp
		fmt.Println("IP Address", temp.IPAddress, "Class Name", temp.ClassName, "Last Octect", temp.LastIPOctet)
		// Lood next Temp Value into temp
		temp = temp.next
	}

	//Print Blank Line
	fmt.Println()
}

// The Main Function Where Things Happen
func main() {

	//CREATE 3 Nodes
	node1 := new(Node)
	node1.IPAddress = "10.10.10.1"
	node1.ClassName = "CLASS A"
	node1.LastIPOctet = 1

	node2 := &Node{"10.10.10.2", "CLASS A", 2, nil}

	node3 := &Node{"10.10.10.3", "CLASS A", 3, nil}

	node4 := &Node{"10.10.10.4", "CLASS A", 4, nil}

	node5 := &Node{"10.10.10.5", "CLASS A", 5, nil}

	// Create the New Stack
	stack := new(Stack)
	// Start Adding Numbers to it
	stack.Push(node1)
	stack.Push(node2)
	stack.Push(node3)
	stack.Push(node4)
	stack.Push(node5)
	// As long as Stack is Empty
	for stack.IsEmpty() == false {
		// Remove from Stack
		val, isValueThere := stack.Pop()
		// Print Results
		fmt.Println("The IP is ", val, " True or False  ", isValueThere)
	}
}
