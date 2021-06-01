package main

import "fmt"

func main() {
	scores := map[string]int{
		"Andres":  1,
		"Labrija": 5,
	}

	fmt.Println(scores)

	scores["Maluma"] = -1
	fmt.Println(scores)

	delete(scores, "Maluma")
	fmt.Println(scores)

	appendisitix(scores)
	fmt.Println(scores)

	var scores2 map[string]int
	scores3 := make(map[string]int)
	scores4 := map[string]int{}

	fmt.Println(scores2["a"])
	fmt.Println(scores3["a"])
	fmt.Println(scores4["a"])

}

func appendisitix(m map[string]int) {
	m["apendisitix"] = 100
}
