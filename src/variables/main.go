package main

import (
	"fmt"
)

var number int

const my_var string = "this is a string" //can't be changed

func main() {
	fmt.Println("variable testing")
	numberr := 10 //short declaration
	janice := 17
	fmt.Println(numberr)
	fmt.Println(my_var)
	fmt.Println(janice)

	hi()
}

func hi() {
	fmt.Printf("eijdiwerjfw %v firwejforghjurtrwge %v", number, my_var)
	fmt.Print("hi")
	fmt.Print("?")
}

// basic type
// integer
// float
//boolean
// aggregate type
// array
// reference type
// interface type
