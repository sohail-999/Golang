package main

import (
	"fmt"
)

func main() {
	a := 4
	squarev(a)
	sqaureadd(&a)
}

func squarev(v int) {
	v *= v

	fmt.Println(v, &v)

}
func sqaureadd(p *int) {

	*p *= *p
	fmt.Println(p, *p)
}
