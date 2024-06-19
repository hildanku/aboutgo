package main

import "fmt"

func modifyValue(x *int) *int {
	*x = 10
	fmt.Printf("Inside function: x = %d\n", *x)
	return x
}

func main() {
	a := 5
	fmt.Printf("Before function call: a = %d\n", a)
	fmt.Println(&a)
	modifyValue(&a)
	fmt.Printf("After function call: a = %d\n", a)
	fmt.Println(&a)
}
