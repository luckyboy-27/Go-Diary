package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func sort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
	}
	return arr
}

func main() {
	arr := []int{23, 431, 65, 353, 653, 76, 343, 76, 1}
	fmt.Println(sort(arr))
}
