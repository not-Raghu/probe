package main

// func main() {

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

// s := []int{1, 2, 3, 4, 5, 6}
// 	//behind the scenes something like
// 	arr := [3]int{10, 20, 30}
// s := arr[:] // slice of the whole array
// fmt.Println(s)

// }

// func main() {
// 	var s []int
// 	fmt.Println(s, len(s), cap(s))
// 	if s == nil {
// 		fmt.Println("nil")
// 	}
// }

// using the make function to make slices
// func main() {

// 	//make([]type,length,capacity);
// 	a := make([]int, 5)
// 	fmt.Println(a)

// 	// for i := 0; i < 10; i++ {
// 	// 	a[i] = i
// 	// }

// 	fmt.Println(a)
// }

// func main() {
// 	s := make([]int, 5, 10)

// 	for i := range 12 {
// 		s[i] = i
// 	}
// //length is the size of slice and capacity is the size from starting point of slice till underlying array (of that slice );
// }

// func main() {
// 	var s []int
// 	s = append(s, 1)
// 	s = append(s, 3)
// 	fmt.Println(s)
// 	s = append(s, 1, 23, 3, 4, 5, 5)
// 	fmt.Println(s)
// }

// func main() {

// 	var nums []int
// 	if nums == nil {
// 		fmt.Println(nil)
// 	}
// }

// func main() {
// 	var nums = make([]int, 3, 4)
// 	var nums2 = make([]int, 3, 4)
// 	copy(nums, nums2)
// 	fmt.Print(nums)
// }
