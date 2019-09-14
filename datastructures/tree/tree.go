package main

import "fmt"

//Class A 0.0.0.0	to 127.255.255.255
//Class B 128.0.0.0	 to 191.255.255.255
// Class C 192.0.0.0	to 223.255.255.255

//type Node struct {
//	Class       string // Class A
//	Network     string // 10.10.10.0
//	FourthOctet string // 128
//	left        *Node
//	right       *Node
//}

type BFSQueue struct {
	start, end *BFSnode
	length     int
}
type BFSnode struct {
	value interface{}
	next  *BFSnode
}

// Create a new queue
func NewQuue() *BFSQueue {
	return &BFSQueue{nil, nil, 0}
}

// Take the next item off the front of the queue
func (this *BFSQueue) Dequeue() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.value
}

// Put an item on the end of a queue
func (this *BFSQueue) Enqueue(value interface{}) {
	n := &BFSnode{value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}

// Return the number of items in the queue
func (this *BFSQueue) Len() int {
	return this.length
}

// Return the first item in the queue without removing it
func (this *BFSQueue) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.start.value
}

type BDFStack struct {
	top    *BDFSNode
	length int
}
type BDFSNode struct {
	value interface{}
	prev  *BDFSNode
}

// Create a new stack
func NewStack() *BDFStack {
	return &BDFStack{nil, 0}
}

// Return the number of items in the stack
func (this *BDFStack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *BDFStack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *BDFStack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *BDFStack) Push(value interface{}) {
	n := &BDFSNode{value, this.top}
	this.top = n
	this.length++
}

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

//Create a Queue
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Funtion to Print nodes as a Level
func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := new(Node)
	curr.value = arr[start]

	left := 2*start + 1  // Odd Nodes to the Left based on Start
	right := 2*start + 2 // Even Node to the Right  based on Start

	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func addValueToNode(n *Node, value int) *Node {
	// There is No Node
	if n == nil {
		n = new(Node)
		n.value = value
		return n
	}
	//  There is a Node
	if value < n.value {
		// Add to the Left
		n.left = addValueToNode(n.left, value)
	} else {
		n.right = addValueToNode(n.right, value)
	}
	return n
}

func (t *Tree) AddNode(value int) {
	t.root = addValueToNode(t.root, value)
}

// Function to Print Pre-Order

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value, "")
	printPreOrder(n.left)
	printPreOrder(n.right)
}
func (t *Tree) PrintPreOrder() {
	fmt.Println("Pre Order ")
	printPreOrder(t.root)
	fmt.Println()
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}

	printInOrder(n.left)
	fmt.Println(n.value, "")
	printInOrder(n.right)
}

func (t *Tree) PrintInOrder() {
	fmt.Println(" Print In Order")
	printInOrder(t.root)
	fmt.Println()
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Println(n.value, "")
}

func (t *Tree) PrintPostOrder() {
	fmt.Println(" Print Post Order")
	printPostOrder(t.root)
	fmt.Println()
}

func (t *Tree) PrintBreadthFirst() {
	que := new(BFSQueue)
	var temp *Node

	if t.root != nil {
		que.Enqueue(t.root)
	}
	fmt.Print("Breadth First : ")
	for que.Len() != 0 {
		temp2 := que.Dequeue()
		temp = temp2.(*Node)

		fmt.Print(temp.value, " ")

		if temp.left != nil {
			que.Enqueue(temp.left)
		}
		if temp.right != nil {
			que.Enqueue(temp.right)
		}
	}
	fmt.Println()
}

func (t *Tree) PrintDepthFirst() {
	stk := new(BDFStack)

	if t.root != nil {
		stk.Push(t.root)
	}
	fmt.Print("Depth First : ")
	for stk.Len() != 0 {
		temp := stk.Pop().(*Node)
		fmt.Print(temp.value, " ")
		if temp.right != nil {
			stk.Push(temp.right)
		}
		if temp.left != nil {
			stk.Push(temp.left)
		}
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Add Nodes to the Tree
	t := LevelOrderBinaryTree(arr)
	t.PrintPreOrder()
	t.PrintPostOrder()
	t.PrintInOrder()
	t.PrintBreadthFirst()
	t.PrintDepthFirst()

}
