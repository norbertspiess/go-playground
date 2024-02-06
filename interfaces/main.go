package main

import "fmt"

type aBot struct{}
type bBot struct{}

type bot interface {
	greet() string
}

func main() {
	a := aBot{}
	b := bBot{}

	printGreet(a)
	printGreet(b)
}

func printGreet(b bot) { // interfaces are implicitly implemented. everything that has methods defined as in interface, fits
	fmt.Println(b.greet())
}

func (aBot) greet() string {
	return "Hey"
}

func (bBot) greet() string {
	return "Hola"
}
