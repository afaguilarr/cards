package deck

import "fmt"

// Gets all the possible card values
func getCardValues() []string {
	return []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"J",
		"Q",
		"K",
	}
}

// Gets all the possible card suits
func getCardSuits() []string {
	return []string{
		"Diamonds",
		"Spades",
		"Clubs",
		"Hearts",
	}
}

// Generates a card based on its value and suit
func getCard(value string, suit string) string {
	return fmt.Sprintf("%s of %s", value, suit)
}
