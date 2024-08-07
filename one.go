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

	//arrays
	var arr [5]int
	fmt.Println(len(arr))
	arr1 := [...]int{1, 2, 3, 4}
	fmt.Println("Element of arr: ", arr[2])
	fmt.Println("Element of arr1: ", arr1[3])

	twoDArr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(twoDArr)

}
