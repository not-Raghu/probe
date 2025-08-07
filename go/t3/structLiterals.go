package main

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

// func main() {
// 	fmt.Println(v1, p, v2, v3) //*p prints the value
// }

// func main() {
// 	v := Vertex{Y: 10, X: 20}
// 	fmt.Println(v)
// }

// var (
// 	pow = []int{1, 2, 3}
// )

// func main() {
// 	fmt.Println(pow)
// }
