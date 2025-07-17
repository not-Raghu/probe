package main

import (
	"fmt"
	"math"
)

// func pow(x, n, lim float64) float64 {
// 	v := math.Pow(x, n)
// 	if v < lim {
// 		return v
// 	}
// 	return lim
// }

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //this is a way to declare a variable inside the condition and it
		return v //this variable is also accessible from else conditions too!
	} else {
		fmt.Print("length is more than the limit `")
		return v
	}
}

// func main() {
// 	fmt.Print(pow(3, 2, 10))
// 	fmt.Println()
// 	fmt.Print(pow(3, 3, 20))
// }
