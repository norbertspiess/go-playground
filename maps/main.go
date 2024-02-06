package main

import "fmt"

func main() {
	colours := map[string]string{
		"white": "#ffffff",
		"black": "#000000",
	}
	colours["red"] = "#ff0000"
	printMap(colours)

	delete(colours, "red")
	printMap(colours)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println(key, ":", val)
	}
}
