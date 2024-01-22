package main

import "fmt"

func check19(i int, arr []int) bool {
	for k := 0; k < len(arr); k++ {
		if arr[k] == i {
			return false
		}
	}
	return true
}

func checkHo(grid [9][9]int, a int) []int {
	arr := []int{}

	for i := 0; i < 9; i++ {
		arr = append(arr, grid[i][a])
	}

	arrH := []int{}

	for i := 1; i <= 9; i++ {
		if check19(i, arr) == true {
			arrH = append(arrH, i)
		}
	}
	return arrH
}

func checkV(grid [9][9]int, b int) []int {
	arr := []int{}

	for i := 0; i < 9; i++ {
		arr = append(arr, grid[b][i])
	}

	arrV := []int{}

	for i := 1; i <= 9; i++ {
		if check19(i, arr) == true {
			arrV = append(arrV, i)
		}
	}
	return arrV
}

func checkB(grid [9][9]int, y int, x int) []int {
	arr := []int{}
	arrb := []int{}
	x = (x / 3) * 3
	y = (y / 3) * 3
	for i := y; i < y+3; i++ {
		for j := x; j < x+3; j++ {
			arr = append(arr, grid[i][j])
		}
	}
	for i := 1; i <= 9; i++ {
		if check19(i, arr) == true {
			arrb = append(arrb, i)
		}
	}
	return arrb
}

func solve(grid [9][9]int) [9][9]int {
	if !isValid(grid) {
		return grid
	}

	if isSolved(grid) {
		return grid
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				h := checkHo(grid, j)
				v := checkV(grid, i)
				b := checkB(grid, i, j)

				s := intersection(h, v)
				w := intersection(s, b)

				for _, candidate := range w {
					grid[i][j] = candidate
					result := solve(grid)
					if isSolved(result) {
						return result
					}
					grid[i][j] = 0
				}
				return grid
			}
		}
	}

	return grid
}

func isValid(grid [9][9]int) bool {
	// Validate the Sudoku
	return true
}

func isSolved(grid [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func intersection(arr1, arr2 []int) []int {
	result := []int{}
	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 {
				result = append(result, val1)
			}
		}
	}
	return result
}

func main() {
	arg := [9]string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}

	grid := [9][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if arg[i][j] == '.' {
				grid[i][j] = 0
			} else {
				grid[i][j] = int(arg[i][j] - '0')
			}
		}
	}

	fmt.Printf("%v", solve(grid))
	fmt.Println()
}
