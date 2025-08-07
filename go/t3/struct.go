package main

// type Employee struct {
// 	id     int
// 	name   string
// 	isPaid bool
// 	age    int
// }

// func main() {
// 	e1 := Employee{1, "raghu", false, 18}
// 	// fmt.Println(e1.age)
// 	// fmt.Println(e1)
// 	e1.isPaid = true
// 	fmt.Println(e1.isPaid)
// }

// func main() {
//the fields are not mandatory , but by default these are given zero values
// e1 := Employee{age: 40}
// fmt.Println(e1)

// if(e1.isPaid){

//

// e1 := Employee{1, "raghu", false, 18}
// p1 := &e1
// p1.name = "not-raghu"
// fmt.Println(p1)

// }

type human struct {
	name  string
	age   int
	class int
}

// func main() {
// 	h1 := human{"raghu", 18, 14}
// 	fmt.Println(h1.name)
// 	fmt.Println(h1.age)
// 	fmt.Println(h1.class)
// }
