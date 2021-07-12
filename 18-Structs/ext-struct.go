package main

import "fmt"

type people struct {
	name string
	age  int
}

type superman struct {
	level int
	people
}

func (p *people) say() {
	fmt.Printf("Im %s,Im normal", p.name)
}

func (s *superman) fly() {
	fmt.Println("I can fly")
}

func main() {

	s := superman{
		level: 1,
		people: people{
			name: "James",
			age:  19,
		},
	}
	s.say()
}
