package main

import "fmt"

/*
thirdtimeisacharm
Instructions
Write a function ThirdTimeIsACharm() that takes a string as an argument and returns another string with every third character.

Return the output followed by a newline \n.
If the string is empty, return a newline \n.
If there is no third character, return a newline \n.
Expected function
*/
func ThirdTimeIsACharm(str string) string {

	if len(str) < 3 {

		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i += 3 {

		result += string(str[i])
	}
	return result
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
