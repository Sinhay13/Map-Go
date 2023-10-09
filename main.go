package main

import "fmt"

func main() {

	//first way to declare a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// second way map
	// var colors map[string]string

	// third way map
	// colors := make(map[string]string)

	// add element
	// colors["white"] = "#ffffff"

	// delete element

	// delete(colors, "white")

	printMap(colors)

	//fmt.Println(colors)

}

// to navigate inside map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}

// for map all the key need to be of same type, the same for the type of the key value

// in map keys are indexed

// map is a reference Type ! no need of pointer no copy of date but copy of reference
