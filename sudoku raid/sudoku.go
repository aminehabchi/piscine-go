package main

import (
	"os"

	"github.com/01-edu/z01"
)

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

func compare(arr1, arr2 []int) []int {
	result := []int{}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				result = append(result, arr1[i])
			}
		}
	}
	return result
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

func solve(grid [9][9]int) [9][9]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				h := checkHo(grid, j)
				v := checkV(grid, i)
				b := checkB(grid, i, j)
				s := compare(h, v)
				w := compare(s, b)
				if len(w) == 1 {
					grid[i][j] = w[0]
				} else {
					for _, r := range w {
						grid[i][j] = r
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
	}

	if isSolved(grid) == false {
		solve(grid)
	}

	return grid
}

func main() {
	arg := os.Args[1:]

	if len(arg) != 9 {

		print("Error\n")
		return
	}

	for i := 0; i < len(arg); i++ {
		if len(arg[i]) != 9 {
			print("Error\n")
			return
		}
	}

	for i := 1; i < len(arg); i++ {
		for j := 0; j < 9; j++ {
			if (arg[i][j] < '0' || arg[i][j] > '9') && arg[i][j] != '.' {
				print("Error\n")
				return
			}
		}
	}
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

	a := solve(grid)

	if !isSolved(a) {

		print("Error\n")
		return
	}
	f := 0
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			f = f + a[j][i]
		}
	}
	if f != 405 {
		print("Error\n")
		return
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune(a[i][j] + '0'))
			if j != 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
