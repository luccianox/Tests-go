/*expandstr
Instructions

Write a program that takes a string and displays it with exactly three spaces between each word,
with no spaces nor tabs at neither the beginning nor the end.

The string will be followed by a newline ('\n').

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := []rune(os.Args[1])
	// loop through the string
	for i := 0; i < len(args); i++ {
		// skip leading spaces
		if i < len(args) && (args[i] == ' ' || args[i] == '\t') {
			i++
		}
		// print a word
		if i < len(args) && (args[i] != ' ' || args[i] != '\t') {
			z01.PrintRune(args[i])
			i++
		}
		// skip spaces after word
		if i < len(args) && (args[i] == ' ' || args[i] == '\t') {
			i++
		}
		// if there are still more words print exactly 3 spaces
		if i < len(args) {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
