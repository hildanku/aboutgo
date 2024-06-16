package main

import "fmt"

func main() {

	var IDProduk int
	var DurasiPembelian uint64
	var TotalHarga uint64

	fmt.Print("Input IDProduk : ")
	fmt.Scan(&IDProduk)

	//AmbilProduk(IDProduk)

	NamaProduk, HargaPerBulan := AmbilProduk(IDProduk)

	fmt.Print("Input Durasi Produk yang ingin dibeli: ")
	fmt.Scan(&DurasiPembelian)

	TotalHarga = uint64(HargaPerBulan) * uint64(DurasiPembelian)

	fmt.Printf("Nama Produk : %s \n", NamaProduk)
	fmt.Print("Total Harga : ", TotalHarga)
}

func AmbilProduk(IDProduk int) (string, uint64) {
	var NamaProduk string
	var HargaPerBulan uint64
	if IDProduk == 1 {
		NamaProduk = "GoLand by JetBrains"
		HargaPerBulan = 720000
	}
	return NamaProduk, HargaPerBulan
}
