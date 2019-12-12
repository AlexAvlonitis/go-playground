package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#dd0000",
		"blue":  "#cc0000",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for _, hex := range c {
		fmt.Println(hex)
	}
}
