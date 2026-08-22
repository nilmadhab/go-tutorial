package main

import (
	"fmt"
)

func squareAdd(p int) {
	p *= p
	fmt.Println("Inside squareAdd, before modification:", p)
}

func squareModify(p *int) {
	(*p) *= (*p)
	fmt.Println("Inside squareAdd, before modification:", *p)
}

func main() {
	// Create a slice of numbers to sum
	//fmt.Println("=== PART 5: POINTERS ===")

	i := 42

	p := &i
	fmt.Printf("Value of i: %d, Address of i: %p\n", i, p)
	// fmt.Println(*p)

	// *p = 100
	// fmt.Printf("New value of i: %d\n", i)

	squareAdd(i)
	fmt.Printf("Value of i after squareAdd: %d\n", i)

	squareModify(&i)
	fmt.Printf("Value of i after squareAdd: %d\n", i)

}
