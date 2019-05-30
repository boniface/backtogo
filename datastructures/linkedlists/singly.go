package main

import "fmt"

type Node struct {
	next *Node
	data interface{}
}

type List struct {
	head *Node
	size int
}

func (list *List) Size() int {
	return list.size
}

func (list *List) isEmpty() bool {
	return list.size == 0
}

func (list *List) Addhead(value interface{}) {
	list.head = &Node{list.head, value}
	list.size++
}

func (list *List) addTail(value int) {
	curr := list.head
	newNode := &Node{nil, value}

	if curr == nil {
		list.head = newNode
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode

}

func (list *List) SortedInsert(value int) {
	newNode := &Node{nil, value}
	curr := list.head

	if curr == nil || curr.data.(int) < value {
		newNode.next = list.head
		list.head = newNode
		return
	}

}

func (list *List) Print() {
	temp := list.head

	for temp != nil {
		fmt.Println("The Data is ", temp.data)
		temp = temp.next
	}
	fmt.Println(" This is a Test ")

}

func main() {

	fmt.Println(" Adding to the Head")
	lst := new(List)

	lst.Addhead(1)
	lst.Addhead(2)
	lst.Addhead(3)
	lst.Addhead(4)
	lst.Print()

	tail := &List{}

	tail.addTail(1)
	tail.addTail(2)
	tail.addTail(3)
	tail.addTail(4)
	tail.Print()

}
