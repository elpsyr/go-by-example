package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40}) //结构体本身并没有代表其自身的内存地址

	ann := person{name: "Ann", age: 40}
	fmt.Println(&ann.name)   //结构体内部的数据有代表其自身的内存地址=
	fmt.Println((&ann).name) //自动解构

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
