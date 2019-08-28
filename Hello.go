package main

import "fmt"

func main() {
	data := []int{9, 3, 6, 7, 8, 2, 1, 4}
	BubbleSort(data, more)
	fmt.Println(" the Sorted Data is ", data)

}

func less(value1 int, value2 int) bool {
	return value1 < value2
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				//swapping
				arr[j+1], arr[j] = arr[j], arr[j+1]

			}
		}
	}
}
