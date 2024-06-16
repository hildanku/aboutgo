package main

import "fmt"

func main() {
	fmt.Println(Produk())
}

func Produk() {
	var NamaProduk string
	var HargaProduk uint16
	NamaProduk = "GoLand"
	HargaProduk = 500000

	fmt.Println(NamaProduk)
	fmt.Println(HargaProduk)
}
