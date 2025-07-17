package main

func Sqrt(x float64) float64 {
	z := 1.0
	for range 10 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// func main() {
// 	// fmt.Scan(i) ig used for taking input
// 	ans := Sqrt(123)
// 	fmt.Print(ans)
// }
