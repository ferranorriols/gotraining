package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var emptyp person
	fmt.Printf("%T %v\n", emptyp, emptyp) //main.person { 0}

	p := person {
		name: "Ferran",
		age: 30,
	}
	fmt.Printf("%T %v\n", p, p) //main.person {Ferran 30}
	fmt.Printf("%T %v\n", p.age, p.age) //int 30
}
