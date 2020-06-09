// this is more like a class deck which extends the slice functionality.
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a deck which has all the functionality of slices/arrays

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, card := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+card)
		}
	}

	return cards
}

// declaring a function as refering to an object d created in the main.go
// then, we give a name to our function here as print and write our code referencing to the reciever(d).

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// function to create a handfull of cards

func deal(d deck, handleSize int) (deck, deck) {
	return d[:handleSize], d[handleSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleDeck() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		rn := r.Intn(len(d))
		d[i], d[rn] = d[rn], d[i]
	}
}
