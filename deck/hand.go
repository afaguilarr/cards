package deck

import "log"

type Hand Deck

// Saves the received deck in the data/deck file
func (h Hand) SaveToFile(env_map map[string]string) {
	SaveToFile(Deck(h), env_map["HAND_FILE"])
}

// Loads the information in the deck/data/hand file and generates a new hand
func LoadHandFromFile(env_map map[string]string) Hand {
	hand, err := LoadFromFile(env_map["HAND_FILE"])
	if err != nil {
		log.Fatal("Error loading hand: ", err)
	}
	return Hand(hand)
}
