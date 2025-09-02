// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/hijack", hijack)
// 	http.ListenAndServe(":3000", nil)
// }

// func hijack(w http.ResponseWriter, r *http.Request) {
// 	hj, ok := w.(http.Hijacker)

// 	if !ok {
// 		fmt.Print("hijaking not possible ")
// 		return
// 	}

// 	conn, buffer, err := hj.Hijack()
// 	if err != nil {
// 		fmt.Print("error", err)
// 		return
// 	}

// 	defer conn.Close()

// 	fmt.Print(&buffer)

// 	// bytes := make([]byte, 1024)
// 	// fmt.Print(conn.Read(bytes))
// 	fmt.Print(buffer.ReadFrom(conn))
// }

// // func not_closure() int {
// // 	count := 0

// // 	count++
// // 	return count
// // }

// // func closure() func() int {
// // 	count := 0
// // 	return func() int {
// // 		count++
// // 		return count
// // 	}
// // }

// // func main() {
// // 	closure := closure()

// // 	closure()
// // 	closure()
// // 	value := closure()
// // 	fmt.Println(value)

// // 	count := not_closure()
// // 	fmt.Print(count)
// // }
