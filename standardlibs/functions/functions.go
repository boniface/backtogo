package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	fmt.Println(" Hello World ", hypot(3, 4))

}
