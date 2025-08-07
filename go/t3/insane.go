package main

type Pointer struct {
	x int
	z uint
}

// func main() {
// 	t := Pointer{1, 4}
// 	p := &t
// 	p.x = 9223372036854775807  // go can store this big number(signed) , 64 bit
// 	p.z = 18446744073709551615 //unsigned 64 bit
// 	fmt.Println(p)
// }
