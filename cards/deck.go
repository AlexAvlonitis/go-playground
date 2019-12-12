package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	for _, face := range cardFaces() {
		for _, number := range cardNumbers() {
			cards = append(cards, number+" of "+face)
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	cardsStringSlice, err := readFile(filename)

	continueUnlessError(err)

	return deck(cardsStringSlice)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func continueUnlessError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	cardsCount := len(d) - 1

	for i := range d {
		randomNum := r.Intn(cardsCount)

		d[i], d[randomNum] = d[randomNum], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) ([]string, error) {
	fileByteSlice, error := ioutil.ReadFile(filename)
	fileStringSlice := strings.Split(string(fileByteSlice), ", ")

	return fileStringSlice, error
}

func cardFaces() []string {
	return []string{"Spades", "Hearts", "Clubs", "Diamonds"}
}

func cardNumbers() []string {
	return []string{"Ace", "Two", "Three", "Four"}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
