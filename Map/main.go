package main

import "fmt"

// func main() {

// 	// var colors map[string]string

// 	// colors := make(map[string]string)

// 	colors := make(map[int]string)
// 	// colors := map[string]string{
// 	// 	"red":   "#ff0000",
// 	// 	"green": "#4bf745",
// 	// }

// 	// fmt.Println(colors)

// 	// colors["white"] = "#ffffff"

// 	colors[10] = "#ffffff"

// 	delete(colors, 10)
// 	fmt.Println(colors)

// }

// iterating over maps
func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// printMap(colors)
	fmt.Println(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
