package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck
//which is slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// d deck syntax is for receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// error handling is tricky beast in go
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	//common pattern in go
	if err != nil {
		// Option #1 - log the error and return call to newDeck()
		// Option #2 - log the error and entirely quit the program -- extreme code issue so must use this one
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//conventions in go is short variables (need to get on with it for now)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//random number generation....
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
