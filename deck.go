package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

//create a new type of deck
//which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := [] string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [] string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	ss := []string(d)
	s := strings.Join(ss, ",")
	return s
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666 )
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")

	return deck(ss)
}