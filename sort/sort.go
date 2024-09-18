package main

import "fmt"

func main() {
	x := []int{3, 4, 2, 5, 9, 14, 3}
	fmt.Println(bubbleSort(x))
	fmt.Println(x)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp

			}
		}
	}
	return arr
}
