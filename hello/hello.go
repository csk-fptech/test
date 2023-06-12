package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	msg, message := greetings.Hello("wo")
	fmt.Println(msg, message)

	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	// 	fmt.Println(a)
	// 	a[i], a[j] = a[j], a[i]
	// 	fmt.Println(a)
	// }

}
