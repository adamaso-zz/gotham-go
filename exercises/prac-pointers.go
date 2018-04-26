package main

import "fmt"

// Fix this
// func main() {
// 	x := "George"
// 	printValue(x)
// 	fmt.Println(x)
// }

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)
}

func printValue(s *string) {
	// print the pointer value
	fmt.Println(s)

	// print the string value
	fmt.Println(*s)

	// change the string value
	*s = "Woop"
}
