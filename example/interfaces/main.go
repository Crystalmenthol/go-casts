package main

import "fmt"

type bots interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {

	rectangle := Rectangle{5, 10}
	square := Square{5}
	circle := Circle{5}

	shapes := []Shaper{rectangle, square, circle}
	for _, shape := range shapes {
		computeArea(shape)
	}

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bots) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// 	VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
