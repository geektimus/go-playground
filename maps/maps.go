package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key,value := range c {
		fmt.Printf("The hex code for %v is %v\n", key, value)
	}
}