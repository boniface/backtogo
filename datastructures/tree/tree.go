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
// Start with A Queue Structure
type BFSQueue struct {
	start, end *BFSnode
	length     int
}
type BFSnode struct {
	value interface{}
	next  *BFSnode
}

// Take the next item off the front of the queue
func (this *BFSQueue) Dequeue() interface{} {
	if this.length == 0 {
		return nil
	}
	node := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return node.value
}

// Put an item on the end of a queue
func (queue *BFSQueue) Enqueue(value interface{}) {
	node := &BFSnode{value, nil}
	if queue.length == 0 {
		queue.start = node
		queue.end = node
	} else {
		queue.end.next = node
		queue.end = node
	}
	queue.length++
}

// Return the number of items in the queue
func (queue *BFSQueue) Len() int {
	return queue.length
}

// Return the first item in the queue without removing it
func (queue *BFSQueue) Peek() interface{} {
	if queue.length == 0 {
		return nil
	}
	return queue.start.value
}

// STACK FOR DEEP FIRST TRAVERSERSAL

type BDFStack struct {
	top    *BDFSNode
	length int
}
type BDFSNode struct {
	value interface{}
	prev  *BDFSNode
}

// Return the number of items in the stack
func (stack *BDFStack) Len() int {
	return stack.length
}

// View the top item on the stack
func (stack *BDFStack) Peek() interface{} {
	if stack.length == 0 {
		return nil
	}
	return stack.top.value
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

// TREE CODE STARTS HERE
type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

// Funtion to Print nodes as a Level
func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	newNode := new(Node)
	newNode.value = arr[start]

	left := 2*start + 1  // Odd Nodes to the Left based on Start
	right := 2*start + 2 // Even Node to the Right  based on Start

	if left < size {
		newNode.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		newNode.right = levelOrderBinaryTree(arr, right, size)
	}
	return newNode
}

func addValueToNode(node *Node, value int) *Node {
	// There is No Node
	if node == nil {
		node = new(Node)
		node.value = value
		return node
	}
	//  There is a Node
	if value < node.value {
		// Add to the Left
		node.left = addValueToNode(node.left, value)
	} else {
		node.right = addValueToNode(node.right, value)
	}
	return node
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
	fmt.Println()
	printPostOrder(t.root)

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
