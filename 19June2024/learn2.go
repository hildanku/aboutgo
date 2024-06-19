package main

import "fmt"

func main() {
	var num int = 2
	var pointerToNum *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Value of pointerToNum:", pointerToNum)
	fmt.Println("Value of num by dereferencing pointerToNum:", *pointerToNum)
}

/* Output
Value of num: 2
Value of pointerToNum: 0xc000112068
Value of num by dereferencing pointerToNum: 2
*/
