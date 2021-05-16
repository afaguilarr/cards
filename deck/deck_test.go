package deck

import (
	// "reflect"
	"strings"
	"testing"
)

// func TestNewDeckWeirdStuff() {
// d := NewDeck()
// sd := NewDeck() Just to test the DeepEqual

// if reflect.TypeOf(d).String() != "deck.Deck" {
//     t.Errorf("Expected deck type to be 'Deck', but got '%s'", reflect.TypeOf(d).String())
// }

// if !reflect.DeepEqual(d, sd) {
//     t.Errorf("\nActual Deck:\n%s\nIs different than the expected deck:\n%s", d.ToString(), sd.ToString())
// }
// }

func TestNewDeckLength(t *testing.T) {
	d := NewDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %s", string(len(d)))
	}
}

func TestNewDeckFormat(t *testing.T) {
	for _, card := range NewDeck() {
		split_card := strings.Split(string(card), " of ")
		suit, value := split_card[0], split_card[1]
		if containsString(getCardSuits(), suit) {
			t.Errorf("The suit '%s' is not an allowed suit", suit)
		}
		if containsString(getCardValues(), value) {
			t.Errorf("The value '%s' is not an allowed value", value)
		}
	}
}

func TestNewDeckUniqueElements(t *testing.T) {
	d := NewDeck()
	if b, s := areAllStringsUnique(d); !b {
		t.Errorf("The card '%s' is not unique", s)
	}
}
