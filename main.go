package main

var counter = 1

func main() {
	cards := newDeckFromFile("cardss.txt")
	cards.print()
}
