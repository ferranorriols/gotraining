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

	//Defining an anonymous struct which contains a previous defined struct
	p2 := struct {
		username string
		person
	} {
		username: "pedro82",
		person: person {
			name: "Pedro",
			age: 34,
		},
	}
	fmt.Printf("%T %v\n", p2, p2)
	fmt.Printf("%T %v\n", &p2, &p2) //*struct { username string; main.person } &{johnsmith {a 34}}
}

