package main

import (
	"fmt"
)

var number int

const my_var string = "this is a string" //can't be changed

func main() {
	fmt.Println("variable testing")
	number := 10
	fmt.Println(number)
	fmt.Println(my_var)

	hi()
}

func hi() {
	fmt.Printf("eijdiwerjfw %v firwejforghjurtrwge %v", number, my_var)
}

// basic type
// integer
// float
//boolean
// aggregate type
// array
// reference type
// interface type
