package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func initPerson() Person {
	return Person{Name: "John", Age: 25}
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", p)

	fmt.Println(initPerson())
}
