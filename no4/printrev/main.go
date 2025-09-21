/*printrevcomb
Instructions
Write a program that prints in descending order on a single line all unique combinations of three different digits
so that the first digit is greater than the second and the second is greater than the third.
These combinations are separated by a comma and a space.*/

package main

import "github.com/01-edu/z01"

func main() {

	for i := '9'; i >= '2'; i-- {
		for j := '8'; j >= '1'; j-- {
			for k := '7'; k >= '0'; k-- {

				if i > j && j > k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i != '2' || j != '1' || k != '0' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
