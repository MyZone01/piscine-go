package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 8; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					z01.PrintRune(rune(48 + i))
					z01.PrintRune(rune(48 + j))
					z01.PrintRune(' ')
					z01.PrintRune(rune(48 + k))
					z01.PrintRune(rune(48 + l))
					if i != 9 || j != 8 || k != 9 || l != 9 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}