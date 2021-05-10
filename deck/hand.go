package deck

type Hand Deck

// Saves the received deck in the data/deck file
func (h Hand) SaveToFile() {
	SaveToFile(Deck(h), "hand")
}

// Loads the information in the deck/data/hand file and generates a new hand
func LoadHandFromFile() Hand {
	return Hand(LoadFromFile("hand"))
}
