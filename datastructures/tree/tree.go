package main

//Class A 0.0.0.0	to 127.255.255.255
//Class B 128.0.0.0	 to 191.255.255.255
// Class C 192.0.0.0	to 223.255.255.255

type Node struct {
	Class       string // Class A
	Network     string // 10.10.10.0
	FourthOctet string // 128
	left        *Node
	right       *Node
}

type Tree struct {
	root *Node
}

// Funtion to Print nodes as a Level
func LevelOrderBinaryTree(arr []Node) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []Node, start int, size int) *Node {
	curr := new(Node)

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func main() {

}
