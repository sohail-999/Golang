package main

import "fmt"

func summarise(nums ...int) { //functions accepting a variable number of argumnets called the Variodic  Functions

	total := 0

	for _, n := range nums {

		total += n
	}
	fmt.Printf("Numbers : %v\n", nums)
	fmt.Printf("Count   : %d\n", len(nums))
	fmt.Printf("Total   : %d\n", total)
}

func main() {

	summarise(1, 2, 3)

	summarise(1, 2, 3, 4)

	summarise(0, 2, 3, 4, 5)
	summarise(0)

}
