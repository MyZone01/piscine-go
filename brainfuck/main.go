package main

import (
	"fmt"
	"os"
)

const SIZE = 2048

func main() {
	if len(os.Args) != 2 {
		return
	}
	program := []byte(os.Args[1])
	buffer := [SIZE]byte{}
	cursor := 0
	openedBracket := 0
	i := 0
	N := len(program)
	for i >= 0 && i < N {
		switch program[i] {
		case '>':
			cursor++
		case '<':
			cursor--
		case '+':
			buffer[cursor]++
		case '-':
			buffer[cursor]--
		case '.':
			fmt.Printf("%c", rune(buffer[cursor]))
		case '[':
			openedBracket = 0
			if buffer[cursor] == 0 {
				for i < N && (program[i] != byte(']') || openedBracket > 1) {
					if program[i] == byte('[') {
						openedBracket++
					} else if program[i] == byte(']') {
						openedBracket--
					}
					i++
				}
			}
		case ']':
			openedBracket = 0
			if buffer[cursor] != 0 {
				for i >= 0 && (program[i] != byte('[') || openedBracket > 1) {
					if program[i] == byte(']') {
						openedBracket++
					} else if program[i] == byte('[') {
						openedBracket--
					}
					i--
				}
			}
		}
		i++
	}
}
