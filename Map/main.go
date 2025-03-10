package main

import (
	"fmt"
)

func main() {
	//declaring map
	//1 way to declare map
	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#00ff00",
			"blue":  "#0000ff",
		}
		fmt.Println(colors)//map[blue:#0000ff green:#00ff00 red:#ff0000]
	*/
	/*
		var colors map[string]string

		fmt.Println(colors) //map[]
	*/

	//3 way of declaring map
	colors := make(map[string]string) //amke is a built in method used to create map
	//to intialize the value
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"

	/*
		// fmt.Println(colors)

		//delete map
		delete(colors, "red")
		fmt.Println(colors)
	*/

	//print key value pair
	printMap(colors)

	fmt.Println()
	//update value
	changeMap(colors)
	printMap(colors)

}

// iterate over map
// make a funtion called printMap
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func changeMap(c map[string]string) {
	// c["red"] = "#ff000"
	c["yellow"] = "#hgfgd" //if some value is not there it will add it
}
