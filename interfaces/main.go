package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	greeting string
}
type spanishBot struct {
	greeting string
}

func main() {
	eb := englishBot{"Hi"}
	sb := spanishBot{"Holi"}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string { return eb.greeting }

func (sb spanishBot) getGreeting() string { return sb.greeting }
