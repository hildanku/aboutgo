package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	ptr := &arr[0]
	fmt.Println(*ptr) // output: 1
	*ptr++
	fmt.Println(*ptr) // output: 2
}
