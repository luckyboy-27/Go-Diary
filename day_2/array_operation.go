package main

import "fmt"

func search(arr []int, value int) int {
	x := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == value {
			x += i
			break
		}
	}
	if x == 0 {
		fmt.Println("Not found!")
	}
	return x
}

func bubbleSort(arr []int) []int {
	var temp int
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func max(arr []int) int {
	temp := arr[0]
	length := len(arr)
	for i := 0; i < length; i++ {
		if temp < arr[i] {
			temp = arr[i]
		}
	}
	return temp
}

func min(arr []int) int {
	temp := arr[0]
	length := len(arr)
	for i := 0; i < length; i++ {
		if temp >= arr[i] {
			temp = arr[i]
		}
	}
	return temp
}

func sum(arr []int) int {
	x := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		x += arr[i]
	}
	return x
}

func removeElement(arr []int, value int) []int {
	length := len(arr)
	new := arr[search(arr, value)+1 : length]
	arr = arr[0:search(arr, value)]
	arr = append(arr, new...)
	return arr
}

//another way
func deleteElement(arr []int, value int) []int {
	length := len(arr)
	pos := search(arr, value)
	for i := pos; i < length-1; i++ {
		arr[i] = arr[i+1]
	}
	arr = arr[0 : length-1]
	return arr
}

func addThings(arr []int, value int, pos int) []int {
	length := len(arr)
	if pos > length {
		pos = length
	}
	for i := length - 1; i >= pos; i-- { //thanks stackoverflow :)))
		arr[i] = arr[i-1]
	}
	arr[pos] = value
	return arr
}

func avg(arr []int) float64 {
	length := len(arr)
	x := 0
	for i := 0; i < length; i++ {
		x += arr[i]
	}
	var mean float64
	mean = float64(x / length)
	return mean
}

func main() {
	arr := []int{1, 3413, 654, 67, 7, 65097, 543, 3655, 5097}
	//fmt.Println(search(arr, 7))
	//fmt.Println(max(arr))
	//fmt.Println(min(arr))
	//fmt.Println(bubbleSort(arr))
	//fmt.Println(sum(arr))
	//fmt.Println(removeElement(arr, 67))
	//fmt.Println(deleteElement(arr, 67))
	//fmt.Println(addThings(arr, 3, 3))
	fmt.Println(avg(arr))
}
