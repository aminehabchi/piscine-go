package piscine

import "github.com/01-edu/z01"

func check(a int, b int, c int, d int, e int, f int, g int, h int) bool {
	board := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	board[0][a] = 1
	board[1][b] = 1
	board[2][c] = 1
	board[3][d] = 1
	board[4][e] = 1
	board[5][f] = 1
	board[6][g] = 1
	board[7][h] = 1
	// vertical
	cas := []int{a, b, c, d, e, f, g, h}
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			if j != i {
				if board[i][cas[j]] == 1 {
					return false
				}
			}
		}
	}
	// x
	// 31524

	for j := 0; j < 8; j++ {
		for i := 1; i < 8; i++ {
			x := j - i
			z := cas[j] - i
			if x < 8 && z < 8 && x >= 0 && z >= 0 {
				if board[x][z] == 1 {
					return false
				}
			}

		}

		for i := 1; i < 8; i++ {
			x := j + i
			z := cas[j] + i
			if x < 8 && z < 8 && x >= 0 && z >= 0 {
				if board[x][z] == 1 {
					return false
				}
			}

		}

		for i := 1; i < 8; i++ {
			x := j - i
			z := cas[j] + i
			if x < 8 && z < 8 && x >= 0 && z >= 0 {
				if board[x][z] == 1 {
					return false
				}
			}

		}

		for i := 1; i < 8; i++ {
			x := j + i
			z := cas[j] - i
			if x < 8 && z < 8 && x >= 0 && z >= 0 {
				if board[x][z] == 1 {
					return false
				}
			}
		}
	}

	return true
}

func EightQueens() {
	for a := 0; a < 8; a++ {
		for b := 0; b < 8; b++ {
			for c := 0; c < 8; c++ {
				for d := 0; d < 8; d++ {
					for e := 0; e < 8; e++ {
						for f := 0; f < 8; f++ {
							for g := 0; g < 8; g++ {
								for h := 0; h < 8; h++ {
									if check(a, b, c, d, e, f, g, h) == true {
										z01.PrintRune(rune(a + 1 + '0'))
										z01.PrintRune(rune(b + 1 + '0'))
										z01.PrintRune(rune(c + 1 + '0'))
										z01.PrintRune(rune(d + 1 + '0'))
										z01.PrintRune(rune(e + 1 + '0'))
										z01.PrintRune(rune(f + 1 + '0'))
										z01.PrintRune(rune(g + 1 + '0'))
										z01.PrintRune(rune(h + 1 + '0'))
										z01.PrintRune('\n')

									}
								}
							}
						}
					}
				}
			}
		}
	}
}
