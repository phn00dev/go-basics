package main

import (
	"fmt"
)

type person struct {
	firstname string
	age       int
}

type Human struct {
	fullName string
	age      int
	weight   int
	height   int
}

type Student struct {
	Human
	specialty string
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {

	var hudayberdi person

	hudayberdi.age = 24
	hudayberdi.firstname = "Hudayberdi"
	fmt.Println("The person's name is", hudayberdi.firstname)

	var polat person
	polat.age = 18
	polat.firstname = "Polat"
	tb_Older, tb_diff := Older(hudayberdi, polat)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", hudayberdi.firstname, polat.firstname, tb_Older.firstname, tb_diff)

	hudayberdipolat := Student{
		Human: Human{
			fullName: "Hudayberdi Polatov",
			age:      24,
			weight:   68,
			height:   167,
		},
		specialty: "Information technology",
	}

	fmt.Println(hudayberdipolat)
}
