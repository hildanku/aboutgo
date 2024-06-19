package main

import "fmt"

func zero(x *int) *int {
	*x = 5
	return x
}
func main() {
	x := 7
	fmt.Println("before,", x) // x is still 5
	zero(&x)
	fmt.Println("after", x) // x is still 5
}
