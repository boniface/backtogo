package main

type Edge struct {
	destination int
	cost        int
	next        *Edge
}

type Graph struct {
	count int
	Edges [](*Edge)
}

func main() {

}
