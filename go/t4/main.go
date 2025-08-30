package main

import (
	"fmt"
)

// func main() {

// 	router := gin.Default()

// 	router.GET("/", func(c *gin.Context) {
// 		c.String(200, "%s", "hello there")
// 	})

// 	router.Run(":3000")
// }

// func hijack(w http.ResponseWriter, r *http.Request) {
// 	hj := w.(http.Hijacker)

// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	conn, buff, err := hj.Hijack()

// }

func not_closure() int {
	count := 0

	count++
	return count
}

func closure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	closure := closure()

	closure()
	closure()
	value := closure()
	fmt.Println(value)

	count := not_closure()
	fmt.Print(count)
}
