package main

import "fmt"

func checkRowsAndColsFor0s(matrix [][]int) {
	rowsWith0 := make([]int, 0)
	colsWith0 := make([]int, 0)

	numRows := len(matrix)
	numCols := len(matrix[0])

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if matrix[i][j] == 0 {
				rowsWith0 = append(rowsWith0, i)
				colsWith0 = append(colsWith0, j)
			}
		}
	}

	for _, row := range rowsWith0 {
		for i := 0; i < numCols; i++ {
			matrix[row][i] = 0
		}
	}

	for _, col := range colsWith0 {
		for i := 0; i < numRows; i++ {
			matrix[i][col] = 0
		}
	}
}

func main() {
	mat := [][]int{
		{1, 2, 3, 4, 5, 6},
		{0, 2, 7, 6, 5, 4},
		{1, 3, 5, 7, 9, 0},
	}

	checkRowsAndColsFor0s(mat)

	fmt.Printf("'%v'\n", mat)
}
