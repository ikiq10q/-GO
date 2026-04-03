package main

import (
	"fmt"
)

func main() {

	// zadanie 1
	matrix := [3][5]int{}

	for i, row := range matrix {
		for j, _ := range row {
			fmt.Scan(&matrix[i][j])
		}

	}
	fmt.Println(matrix)

	max := matrix[0][0]
	max_i, max_j := 0, 0

	for i, row := range matrix {
		for j, matrix := range row {
			if matrix > max {
				max = matrix
				max_i = i
				max_j = j
			}
		}
	}

	fmt.Printf("максимально число: %d at (%d,%d)\n", max, max_i, max_j)

	//zadanie 2

	// заполняем матрицу
	star := [11][11]string{}
	for i, row := range star {
		for j, _ := range row {
			star[i][j] = "."
		}

	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	// заполняем звездами
	for i, row := range star {
		for j, _ := range row {

			// главная диагональ
			if i == j {
				star[i][j] = "*"
			}

			// побочная диагональ
			if i+j == len(star)-1 {
				star[i][j] = "*"
			}

			// i = средняя строка j = средний столбец
			if i == len(star)/2 || j == len(star)/2 {
				star[i][j] = "*"
			}
			fmt.Print(star[i][j], " ")
		}
		fmt.Println()
	}

	// zadanie 3
	chess := [8][8]string{}

	for i, row := range chess {
		for j, _ := range row {
			if (i+j)%2 == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Println()

		//zadanie 4

		matrix_2 := [4][4]int{}
		for i, row := range matrix_2 {
			for j, _ := range row {
				fmt.Scan(&matrix_2[i][j])
			}

		}

		fmt.Println("Введие какие две строки поменять:")

		var row_1, row_2 int
		fmt.Scan(&row_1, &row_2)

		for _, row := range matrix_2 {
			for _, matrix_2 := range row {
				fmt.Print(matrix_2, " ")
			}
			fmt.Println()
		}

		fmt.Println()

		for i, row := range matrix_2 {
			for j, _ := range row {
				if i == row_1-1 {
					matrix_2[row_1-1][j], matrix_2[row_2-1][j] = matrix_2[row_2-1][j], matrix_2[row_1-1][j]
				}
			}

		}
		for _, row := range matrix_2 {
			for _, matrix_2 := range row {
				fmt.Print(matrix_2, " ")
			}
			fmt.Println()
		}

		//zadanie 5

		matrix_3 := [4][4]int{}
		for i, row := range matrix_3 {
			for j, _ := range row {
				fmt.Scan(&matrix_3[i][j])
			}

		}

		fmt.Println("Введие какие два слобца поменять:")

		var col_1, col_2 int

		fmt.Scan(&col_1, &col_2)

		for _, row := range matrix_3 {
			for _, matrix_3 := range row {
				fmt.Print(matrix_3, " ")
			}
			fmt.Println()
		}

		fmt.Println()

		for i := range matrix_3 {
			matrix_3[i][col_1-1], matrix_3[i][col_2-1] = matrix_3[i][col_2-1], matrix_3[i][col_1-1]
		}

		for _, row := range matrix_3 {
			for _, matrix_3 := range row {
				fmt.Print(matrix_3, " ")
			}
			fmt.Println()
		}
	}

}
