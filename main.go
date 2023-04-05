package main

import "fmt"

func main() {

	card := newCard() // := is used for intilization

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
