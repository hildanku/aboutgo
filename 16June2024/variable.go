package main

import "fmt"

/*
Di Golang terdapat 7 cara untuk mendeklarasikan sebuah variable
1.Deklarasi Variabel Beserta Tipe Data
2.Deklarasi Variabel Menggunakan Keyword var
3.Deklarasi Variabel Tanpa Tipe Data
4.Deklarasi Multi Variabel
5.Variabel Underscore _
6.Deklarasi Variabel Menggunakan Keyword new
7.Deklarasi Variabel Menggunakan Keyword make
*/

func main() {
	DeclarateWithDataType()
	DeclarateWithoutDataType()
	DeclarateMultipleVariable("Hallo, Hildan", "Aku berasal dari fungsi DeclarateMultipleVariable", "!")
	UndescoreVariable()
	DeclarateWithNew()
}

func DeclarateWithDataType() {
	// Deklarasi Variabel beserta tipe data / manifest typing
	var to string
	to = "Hildan"
	var pesan string = "this is variable declaration with data type"
	fmt.Println(to, pesan)
	// variable %s akan merubah kata dengan variable string sesuai parameter selanjutnya
	fmt.Printf("Hi %s, %s \n", to, pesan)

}

func DeclarateWithoutDataType() {
	// golang juga mengadopsi type inference
	// inference yaitu metode deklarasi variable yang tipe data-nya
	// diketahui secara otomatis dari nilai tersebut
	pesan := "Halo Hildan, Saya dari DeclarateWithoutDataType"
	fmt.Println(pesan)
}

func DeclarateMultipleVariable(first string, second string, third string) {
	fmt.Println(first, second, third)
}

func UndescoreVariable() {
	// UnderScore adalah reserved variable yang bisa dimanfaatkan untuk
	// menampung nilai yang tidak terpakai, also called "recycle bin"
	// variable undescore biasannya dimanfaatkan untuk menampung return fungsi yang tidka digunakan
	// varibale undescore tidak dapat ditampilkan, ibarat seperti blackhole
	name, _ := "Haloo", "xixixi"
	fmt.Println(name)
}

func DeclarateWithNew() {
	// Deklarasi menggunakan keyword new
	// new diganakan untuk membuat variable pointer dengan tipe data tertentu

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

	// explore dikit
	*name = "Test"
	fmt.Println(*name)
}
