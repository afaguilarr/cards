package deck

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Deck []string

// Generates a new deck (not shuffled)
func NewDeck() Deck {
	d := Deck{}
	for _, suit := range getCardSuits() {
		for _, value := range getCardValues() {
			d = append(d, getCard(value, suit))
		}
	}
	log.Println("New Deck:", d)
	return d
}

// Returns the string representation of the received deck
func (d Deck) ToString() string {
	s := strings.Join(d, "\n")
	log.Println("Deck String:\n" + s)
	return s
}

// Returns the bytes representation of the received deck
func (d Deck) ToBytes() []byte {
	b := []byte(d.ToString())
	log.Println("Deck Byte Slice:", b)
	return b
}

// Saves the received deck in the data/deck file
func (d Deck) SaveToFile(env_map map[string]string) {
	SaveToFile(d, env_map["DECK_FILE"])
}

func SaveToFile(d Deck, fileName string) {
	err := ioutil.WriteFile(fileName, d.ToBytes(), 0666)
	if err != nil {
		log.Fatal("Error while trying to save a deck: ", err)
	}
	log.Println(fmt.Sprintf("%s successfully saved", fileName))
}

// Loads the information in the data/deck file and generates a new deck
func LoadDeckFromFile(env_map map[string]string) (Deck, error) {
	return LoadFromFile(env_map["DECK_FILE"])
}

func LoadFromFile(fileName string) (Deck, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	deckString := string(file)
	deckStringSlice := strings.Split(deckString, "\n")

	log.Println(fmt.Sprintf("%s String Slice:", fileName), deckStringSlice)
	return deckStringSlice, nil
}

// Gets a hand composed of the defined number of cards and also returns
// the remaining deck
func (d Deck) GetHandAndRemainingDeck(handSize int) (Hand, Deck) {
	hand := d[:handSize]
	deck := d[handSize:]
	log.Println("Hand:", hand)
	log.Println("Remaining Deck:", deck)
	return Hand(hand), deck
}

// Shuffles the received deck
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	log.Println(d)
}
