package main

import "fmt"

// func main() {
// 	m := make(map[string]int)
// 	m["kim"] = 0
// 	m["john"] = 1
// 	m["seon"] = 2

// 	fmt.Println(m)
// 	fmt.Println(len(m))
// 	fmt.Println(m["john"])
// }

// func main() {
// 	var pow = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
// 	for i, v := range pow {
// 		fmt.Printf("%d**%d = %d\n", i, i, v)
// 	}

// }

func main() {
	m := make(map[int]int)
	m[1] = 12
	m[2] = 23
	for i := range m {
		fmt.Println(m[i])
	}
}
