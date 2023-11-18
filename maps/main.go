package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("color:%s, hex:%s \n", color, hex)
	}
}
func main() {
	//var colors map[string]string

	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)

	// delete(colors, "white")

	fmt.Println(colors)
}
