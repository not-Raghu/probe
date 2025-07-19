package main

func compute(fn func(int, int) int, x int, y int) int {
	// first param is a func
	ans := fn(x, y)
	return ans
}

// func main() {
// 	add := func(x int, y int) int {
// 		return x + y
// 	}

// 	// fmt.Print(add(1, 2))
// 	ans := compute(add, 1, 2)
// 	fmt.Print(ans)

// }
