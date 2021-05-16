package main

import (
	"cards/deck"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	setLogsState(os.Args)
	d := deck.NewDeck()
	hand, d := d.GetHandAndRemainingDeck(7)
	hand.SaveToFile()
	d.SaveToFile()
	d = deck.LoadDeckFromFile()
	hand = deck.LoadHandFromFile()
	d.Shuffle()

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range s {
		fmt.Printf("%v is %s\n", num, oddOrEvens(num))
	}
}

func oddOrEvens(num int) string {
	if num%2 != 0 {
		return "odd"
	}
	return "even"
}

func setLogsState(osArgs []string) {
	areLogsEnabled := ""
	if len(osArgs) > 1 {
		areLogsEnabled = osArgs[1]
	}
	if areLogsEnabled == "no-logs" {
		log.SetOutput(ioutil.Discard)
	}
}
