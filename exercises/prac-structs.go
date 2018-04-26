package main

import "fmt"

// declare a struct called Movie
// add the following fields:
// - Title (string)
// - Released (bool)
// - Length (int)

type Movie struct {
	Title string
	Released bool
	Length int
}

func main() {
	// declare a variable called "movie" of type "Movie"
	// Set the Title to "Wizard of Oz"
	// Set the Released variable to "true"
	// Set Length to 125

	movie := Movie {
		Title: "Wizard of Oz",
		Released: true,
		Length: 125,
	}

	// Print the value of "movie" out
	// hint: you can use fmt.Println(movie)

	fmt.Println(movie);
}