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
	fmt.Println(person{age: 35})
	fmt.Println(person{age: 40, name: "Luis"})
	fmt.Println(&person{name: "Ann", age: 50})
	s := person{"Sean", 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(s.age)
}
