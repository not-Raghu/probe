package main

import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	defer fmt.Print(add(1, 2)) //executes when the surrounding function returns
// 	fmt.Print("the value of the sum is: ")
// }

func main() {
	fmt.Println("counting...")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	for i := range 10 {
		defer fmt.Println(i)
	}

	//basically pushed to stack
}
