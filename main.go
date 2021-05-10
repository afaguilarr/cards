package main

import (
    "cards/deck"
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
}

func setLogsState(osArgs []string) {
    areLogsEnabled := ""
    if len(osArgs) > 1 {areLogsEnabled = osArgs[1]}
    if areLogsEnabled == "no-logs" {log.SetOutput(ioutil.Discard)}
}
