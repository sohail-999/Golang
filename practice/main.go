package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	//var name string = "Sohail"
	//var age int = 22
	//var bookingsforsohail = [50]string{}
	//bookingsforsohail[0] = "Room no 21"
	//bookingsforsohail[1] = "Room no 22"
	//bookingsforsohail[2] = "Room no 23"

	//fmt.Println("Here is your bookings", name, ":", "\n", bookingsforsohail[0], "\n", bookingsforsohail[1], "\n", bookingsforsohail[2])

	//fmt.Println("Here is your name :\n", name, "\nHere is your age", name, ":", "\n", age)
	color.Yellow("Hello in color!") // actually using it
}
