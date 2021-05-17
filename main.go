package main

import (
	"cards/deck"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	setLogsState(os.Args)
	env_map := set_env_map()
	d := deck.NewDeck()
	hand, d := d.GetHandAndRemainingDeck(7)
	hand.SaveToFile(env_map)
	d.SaveToFile(env_map)
	d, _ = deck.LoadDeckFromFile(env_map)
	hand = deck.LoadHandFromFile(env_map)
	d.Shuffle()
}

func set_env_map() map[string]string {
	env_map, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	log.Println(env_map)
	return env_map
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
