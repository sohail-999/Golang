package main

import (
	"fmt"
)

func main() {

	count := 0

	var increment func() = func() {

		count++
		fmt.Println(count)

	}

	increment()
	increment()
	increment()

}
