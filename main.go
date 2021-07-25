package main

var counter = 1

func main() {
	cards := newDeck()
	//cards.saveToFile("asd.txt")
	//newCards := newDeckFromFile("afsd.txt")
	cards.shuffle()
	cards.print()
}
