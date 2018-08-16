package main

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	fmt.Println(" Hello World ", hypot(3, 4))
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
	fmt.Printf("%T\n", hypot)
	fmt.Printf("%T\n", methodA)
	fmt.Printf("%T\n", methodB)
	fmt.Printf("%T\n There is a function ", math.Sin)

	data := 10
	ptr := &data
	fmt.Println("Value stored at variable var is ", data)
	fmt.Println("Value stored at variable var is ", *ptr)
	fmt.Println("The address of variable var is ", &data)
	fmt.Println("The address of variable var is ", ptr)

}

func methodA(i, j, k int, s, t string) {

}

func methodB(i int, j int, k int, s string, t string) {

}

func add(x, y int) int {
	return x + y
}

func sub(x int, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}
