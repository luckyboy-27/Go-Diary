package main

import "fmt"

func matrix_add(a [100][100]int, b [100][100]int, row int, col int) int {
	var sum [100][100]int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = a[i][j] + b[i][j]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ", sum[i][j])
			if j == col-1 {
				fmt.Println("")
			}
		}
	}
	return 0
}

func matrix_mul(a [100][100]int, b [100][100]int, row1 int, col1 int, row2 int, col2 int) int {
	if col1 != row2 {
		fmt.Println("the size of matrices is not suitable. Number of Columns in matrix 1 must be equal to number of columns in matrix 2")
	} else {
		var i, j, k int
		var product [100][100]int
		for i = 0; i < row1; i++ {
			for j = 0; j < col2; j++ {
				for k = 0; k < col1; k++ {
					product[i][j] += (a[i][k] * b[k][j])
				}
			}
		}
		for i = 0; i < row1; i++ {
			for j = 0; j < col2; j++ {
				fmt.Printf(" %d ", product[i][j])
				if j == col2-1 {
					fmt.Println("")
				}
			}
		}
	}
	return 0
}

func main() {
	var matrix1 [100][100]int
	var matrix2 [100][100]int
	var row1, col1, row2, col2 int
	fmt.Print("Enter number of rows for matrix 1: ")
	fmt.Scanln(&row1)
	fmt.Print("Enter number of cols for matrix 1: ")
	fmt.Scanln(&col1)

	fmt.Println()
	fmt.Println("Matrix 1:")
	fmt.Println()
	for i := 0; i < row1; i++ {
		for j := 0; j < col1; j++ {
			fmt.Printf("Enter the element for Matrix 1 %d %d :", i+1, j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}
	fmt.Println()
	fmt.Print("Enter number of rows for matrix 2: ")
	fmt.Scanln(&row2)
	fmt.Print("Enter number of cols for matrix 2: ")
	fmt.Scanln(&col2)

	fmt.Println()
	fmt.Println("Matrix 2:")
	fmt.Println()

	for i := 0; i < row1; i++ {
		for j := 0; j < col1; j++ {
			fmt.Printf("Enter the element for Matrix 2 %d %d :", i+1, j+1)
			fmt.Scanln(&matrix2[i][j])
		}
	}
	//fmt.Println()
	//fmt.Println("Sum of Matrix: ")
	//fmt.Println()

	//matrix_add(matrix1, matrix2, row, col)

	fmt.Println()
	fmt.Println("Product of Matrix: ")
	fmt.Println()

	matrix_mul(matrix1, matrix2, row1, col1, row2, col2)

}
