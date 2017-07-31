package main

import "fmt"

type person struct{
	name string
	age int
}

func main(){
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name:"Alice", age:111})

	fmt.Println(person{name:"hello"})

	fmt.Println(person{age:122})

	fmt.Println(&person{name:"klaus", age: 23})

	s := person{name:"jacob", age:100}
	fmt.Println(s.name)

	sp := &s

	sp.age = 101
	fmt.Println(*sp)
	fmt.Println(sp)
	fmt.Println(s)
}
