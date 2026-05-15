package main

import "fmt"

func add(a int, b int) (int, int) { //multiple return types
	// returns TWO ints
	sum := a + b

	product := a * b

	return sum, product
}

func main() {

	s, p := add(5, 5)

	fmt.Println(s) // 10

	fmt.Println(p) // 25

}
