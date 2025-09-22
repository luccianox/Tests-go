/*
itoa
Instructions
Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as
anint in a number represented as a string.

For this exercise the handling of the signs + or - does have to be taken into account.

Expected function
*/
package main

import "fmt"

func Itoa(n int) string {

	if n == 0 {

		return "0"
	}

	sign := ""
	if n < 0 {

		n = -n
		sign = "-"
	}
	result := ""
	for n > 0 {

		nb := n % 10
		result = string(rune('0'+nb)) + result

		n /= 10

	}
	return sign + result

}
func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
