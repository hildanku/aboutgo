package main

import "fmt"

// IMPORTANT
//

func main() {
	fmt.Println("Hi, i'm a data type")
	NumericNonDecimalDataType()
	NumericDecimal(2.1)
	CheckPassword("passwordnyaapahayo")
	CheckPasswordCondition("passwordnyaapahayo")
}

func NumericNonDecimalDataType() {
	// numerik non desimal ada 2 jenis yaitu uint (bilangan cacah atau positif)
	// dan int (bilangan bulat atau positif dan negatif)
	// uint dan int dipecah lagi
	// docs : https://dasarpemrogramangolang.novalagung.com/A-tipe-data.html#a101-tipe-data-numerik-non-desimal
	// Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variabel,
	// sebisa mungkin tipe yang dipilih harus disesuaikan dengan nilainya,
	// karena efeknya adalah ke alokasi memori variabel.
	// Pemilihan tipe data yang tepat akan membuat pemakaian memori lebih optimal, tidak berlebihan.
	var positif uint8 = 255
	var negative = -2147483648

	fmt.Printf("Positif: %d\n", positif)
	fmt.Printf("Negative: %d\n", negative)

	// %d digunakan untuk memformat data numerik non desimal
	// pada variable negative, disini tidak di deklarasikan tipe datannya,
	// compiler secara otomatis akan mendeteksi tipe data dari nilai yang di assign
	// karena value dari negative masuk scope int32
	// maka otomatis tipe data dari variable tersebut adalah int32
}

func NumericDecimal(decimal float32) {
	fmt.Println(decimal)
}

func CheckPassword(Password string) {
	message := "Selamat anda mendapatkan sembako"
	if Password == "passwordnyaapahayo" {
		fmt.Println(message)
	} else {

	}
}
func CheckPasswordCondition(Password string) bool {
	if Password == "passwordnyaapahayo" {
		return true
	} else {
		return false
	}
}
