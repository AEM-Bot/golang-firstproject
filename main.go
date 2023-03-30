package main

// import "fmt"

func main() {

	// fmt.Println("Initial Code for Cards Game!!")
	// var card string = "Ace of Spades"

	//only for declaring new variable and initialization
	// card := "Ace of Spades"

	// card := newCard()

	// card = "Five of Diamonds"

	// fmt.Println(card)

	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := deck{"Ace of Diamods", newCard()}

	// fmt.Println(cards)

	// cards.print()

	// fmt.Println(cards)

	// cards := deck(newDeck())

	// hand, remainingCards := deal(cards, 5)

	// // cards.print()
	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

//array and slice data structure in go --> slice can grow or shrink
//must be of same data type
//Go is not OOPS
