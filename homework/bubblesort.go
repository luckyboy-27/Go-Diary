package main

import "fmt"

func swap(arr []int, i int) {
	temp := arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = temp
}

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				swap(arr, j)
			}
		}
	}
	return arr
}

func main() {
	var num int
	fmt.Printf("Enter the size of your array: ")
	fmt.Scan(&num)
	arr := make([]int, num)
	fmt.Println("Enter elements of array: ")
	for i := 0; i < num; i++ {
		fmt.Printf("Enter element %d: ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println(bubbleSort(arr))
}
