package main

import "fmt"

// Node Structure
type Node struct {
	value int
	next  *Node
}

// The Stuck Structure
type Stack struct {
	head *Node
	size int
}

// Methods push

// Get the Size of the Stack
func (stack *Stack) Size() int {
	return stack.size
}

// Is the Stack Empty

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) Peek() (int, bool) {
	if stack.IsEmpty() {
		fmt.Println(" the Stack is Empty")
		return 0, false
	}
	return stack.head.value, true
}

func (stack *Stack) Push(value int) {
	stack.head = &Node{value, stack.head}
	stack.size++
}

func (stack *Stack) Pop() (int, bool) {
	//Check if Stack is Empty
	if stack.IsEmpty() {
		fmt.Println("Stack is Empty ")
	}
	//Get the Stack value
	value := stack.head.value
	//Update the Head to the next Top Remaining Stack
	stack.head = stack.head.next
	//Reduce the Size of the Stack
	stack.size--

	//Return the Results
	return value, true
}

func (stack *Stack) Print() {
	// Put the Stack Head into the temporary Node
	temp := stack.head
	// Print Value
	fmt.Println("Value put in the Stack ::")
	// If the Temp had NO Error
	for temp != nil {
		// Print the Value of the Node Stored in Temp
		fmt.Println(temp.value, "")
		// Lood next Temp Value into temp
		temp = temp.next
	}

	//Print Blank Line
	fmt.Println()
}

// The Main Function Where Things Happen
func main() {
	// Create the New Stack
	stack := new(Stack)
	// Start Adding Numbers to it
	stack.Push(10)
	stack.Push(5)
	stack.Push(1)
	// As long as Stack is Empty
	for stack.IsEmpty() == false {
		// Remove from Stack
		val, isValueThere := stack.Pop()
		// Print Results
		fmt.Println("The Values are ", val, " True or False  ", isValueThere)
	}
}
