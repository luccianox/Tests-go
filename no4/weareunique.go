/*
weareunique
Instructions
Write a function that takes two strings's and returns the number of characters that are not included in both,
without repeating characters.

If there is no unique characters return 0.
If both strings are empty return -1.
Expected function
*/
package main

import "fmt"

func WeAreUnique(str1, str2 string) int {

	if len(str1) == 0 && len(str2) == 0 {

		return -1
	}

	count := 0

	r1 := []rune(str1)
	r2 := []rune(str2)

	used := []rune{}

	//loop thru str1

	for _, ch := range str1 {
		if !inSlice(r2, ch) && !inSlice(used, ch) {
			count++
			used = append(used, ch)

		}

	}
	for _, ch := range str2 {

		if !inSlice(r1, ch) && !inSlice(used, ch) {

			count++
			used = append(used, ch)
		}
	}
	return count
}
func inSlice(s []rune, r rune) bool {

	for _, v := range s {

		if v == r {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
