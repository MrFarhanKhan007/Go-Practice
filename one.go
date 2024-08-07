package main

import (
	"fmt"
	// "math/rand/v2"
	"reflect"
	"time"
)

// import "rsc.io/quote"

func main() {
	// // fmt.Println("Hello World");
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Hello())

	//strings
	a := "Farhan"
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	//constant
	const s string = "constant"

	fmt.Println(s)

	// for loop
	for i := 0; i < 10; i++ {
		// fmt.Print(i)
	}

	//if-else
	num := 4
	if num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}

	//switch
	// day := rand.IntN(7)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

}
