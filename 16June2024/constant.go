package main

import "fmt"

func main() {
	const Introduction string = "Hi, i'm a constant"
	fmt.Println(Introduction)
	SingleConst("jangan lupa selalu ABCDE")
	MultipleConst()
}

func SingleConst(Input string) {
	const KataKataHariIni string = "jangan lupa selalu ABCDE"
	if Input != KataKataHariIni {
		fmt.Println("Salah!")
	} else {
		fmt.Println("Alhamdulillah")
	}
}

func MultipleConst() {
	const (
		Hari            = "Minggu"
		KataKataHariIni = "Jangan lupa ABCDE\n"
	)
	fmt.Print(KataKataHariIni)
	fmt.Println(Hari)
	// Perbedaan Print dan PrintLn adalah
	// ketika kita menggunakan Print, maka fungsi tersebut tidak otomatis membuatkan line baru
	// sehingga kita perlu \n untuk menambahkan line
	// Ketika kita menggunakan PrintLn, maka secara otomatis akan membuat line baru
}
