package main

import "fmt"

type Person struct {
	id   int
	name string
}

// func (p Person) print() {
// 	fmt.Println("id is", p.id)
// 	fmt.Println("name is", p.name)
// }

// func main() {
// 	p := Person{1, "raghu"}
// 	p.print()
// }

type MyInt int

func (x MyInt) printX() {
	fmt.Print(x)
}

// func main() {
// 	i := MyInt(1)
// 	i.printX()
// }
