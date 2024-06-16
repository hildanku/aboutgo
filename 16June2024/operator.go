package main

import "fmt"

func main() {
	fmt.Println(Calculate(2, 3))
	CalculateFmt(2, 67)
	OddEvenDetector(7)
	Test(2)
	OddEvenDetectorLoop(2)
}

// Kode dengan kesalahan berfikir
func Test(num int) {
	for i := 0; i <= num; i++ {
		if num%2 == 0 {
			fmt.Printf("Angka %d Genap\n", i)
		} else {
			fmt.Printf("Angka %d Ganjil\n", i)
		}
	}
}

// Fungsi dengan kebenaran berfikir
func OddEvenDetectorLoop(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Printf("Angka %d Genap\n", i)
		} else {
			fmt.Printf("Angka %d Ganjil\n", i)
		}
	}
}

func OddEvenDetector(num int) {
	if num%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}

// Tanpa Menggunakan Return
func CalculateFmt(a int, b int) {
	res := a * b
	fmt.Println(res)
}

// Menggunakan Return
func Calculate(a int, b int) int {
	res := a * b
	return res
}
