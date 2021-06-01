package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Using just contactInfo would create a field of name contactInfo with type contactInfo, pretty ugly
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	remove(&slice)
	fmt.Println(slice)
	// pi := contactInfo{"afaguilardo@asdasd.com", 123}
	// p := person{"Andres", "Aguilar", pi}
	// p.updateFirstName("Andy")
	// p.print()
}

func remove(s *[]int) {
	*s = (*s)[:2]
	fmt.Println(*s)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
