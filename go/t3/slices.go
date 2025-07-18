package main

import "fmt"

func main() {

	// primes := [6]int{2, 3, 5, 7, 11, 13}

	// var s []int = primes[1:4] //[low : high]
	// fmt.Println(s)

	//A slice does not store any data, it just describes a section of an underlying array.
	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	//s2 := primes[:] //copy everything
	// fmt.Println(s2)

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	//slice literals are like dynamic arrays

	s := []int{1, 2, 3, 4, 5, 6}
	// 	//behind the scenes something like
	// 	arr := [3]int{10, 20, 30}
	// s := arr[:] // slice of the whole array
	fmt.Println(s)

}
