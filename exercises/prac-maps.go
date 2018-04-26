package main

import "fmt"

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and the value
	for key, value := range beatles {
		fmt.Println(key, value)
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	_, Bob := beatles["Bob"]
	if !Bob{
		fmt.Println("Bob is not a Beatle")
	}

}