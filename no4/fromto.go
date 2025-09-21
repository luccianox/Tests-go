/*
fromto
Instructions
Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.
Expected function
*/
package main

import "fmt"

func FromTo(from int, to int) string {

	result := ""

	if from > 99 || to > 99 || from < 0 || from < 0 {

		return "invalid\n"
	}
	helper := func(n int) string {

		for n < 0 {

			return "0" + string(rune('0'+n))
		}
		return string(rune('0'+n/10)) + string(rune('0'+n%10))
	}

	if from <= to {

		for n := from; n <= to; n++ {

			result += helper(n)
			if n != to {

				result += ", "

			}
		}

	} else {

		for n := from; n >= 0; n-- {

			result += helper(n)

			if n != to {

				result += ", "
			}
		}
	}

	return result + "\n"
}
func main() {
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
