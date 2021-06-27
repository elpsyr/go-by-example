package main

import "fmt"

type exercise interface {
	Run() string
}

func (p *people) Run() string {
	return fmt.Sprintf("%s -%d :Im'running", p.name, p.age)
}

type people struct {
	name string
	age  int
}

func toGym() exercise {
	return &people{name: "tom", age: 18}
}

func main() {
	//获得exercise实例
	e := toGym()
	fmt.Println(e)
	eType, ok := e.(*people)
	fmt.Println(eType, ok)
	run := e.Run()
	fmt.Println(run)
}
