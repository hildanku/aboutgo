package main

import "fmt"

func main() {
	WriteFile(2)
}

func WriteFile(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("num %d \n", i)
	}
}
