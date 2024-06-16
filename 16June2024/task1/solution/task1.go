package main

import "fmt"

func main() {

	/* Ternyata kita tidak perlu menuliskan 1 variable 1 tipe data secara berulang seperti ini:
		var IDProduk int
		var DurasiPembelian uint64
	ternyata kita bisa menyingkat hal tersebut menjadi seperti dibawah
	*/

	// Inisialisasi Tipe Data
	var (
		IDProduk        int
		DurasiPembelian uint64
		TotalHarga      uint64
	)

	/* Perbedaan fmt.Print, Printf, Println adalah
	print: Mencetak teks ke file tanpa menambahkan newline di akhir teks.
	printf: Mencetak teks ke file dengan menambahkan newline di akhir teks.
	println: Mencetak teks ke file dengan menambahkan newline di akhir teks,
			dan setiap baris teks yang dicetak akan secara otomatis diakhiri dengan newline.
	*/

	fmt.Print("Input IDProduk : ")
	/* Scanf, Fscanf, and Sscanf parse the arguments according to a format string,
	analogous to that of Printf. In the text that follows,
	'space' means any Unicode whitespace character except newline.

	*/
	/*
		* Gunakan Scan bila Anda ingin membaca satu input dan menyimpannya ke dalam satu variable.
		* Gunakan Scanf bila Anda ingin membaca beberapa input dan menyimpannya ke dalam beberapa variable,
			dan Anda ingin menentukan format inputnya.
		* Gunakan Scanln bila Anda ingin membaca input sebagai string, termasuk spasi dan newline.
	*/
	/*
		Scan: Selalu gunakan & di depan variable.
		Scanf: Gunakan & di depan variable untuk setiap nilai yang ingin disimpan.
		Scanln: Gunakan & di depan variable untuk menampung string.
	*/
	fmt.Scan(&IDProduk)
	/*
		Analogi kode diatas adalah,
		pertama user memasukan input,
		input tersebut di reference ke addres IDProduk misal 0xc00000a0b8
	*/
	//fmt.Println(&IDProduk)
	//AmbilProduk(IDProduk)

	/*
		Ambil nilai return dari fungsi AmbilProduk berdasarkan IDProduk yang diinput
	*/
	NamaProduk, HargaPerBulan := AmbilProduk(IDProduk)

	fmt.Print("Input Durasi Produk yang ingin dibeli: ")
	fmt.Scan(&DurasiPembelian)

	TotalHarga = HargaPerBulan * DurasiPembelian

	fmt.Printf("Nama Produk : %s \n", NamaProduk)
	fmt.Print("Total Harga : ", TotalHarga)
}

// AmbilProduk membutuhkan parameter IDProduk
// Maksud dari (string, uint64) merujuk kepada inisialisasi variable
// yang difungsikan agar bisa mengembalikan return nilai
func AmbilProduk(IDProduk int) (string, uint64) {
	var NamaProduk string
	var HargaPerBulan uint64
	if IDProduk == 1 {
		NamaProduk = "GoLand by JetBrains"
		HargaPerBulan = 720000
	}
	return NamaProduk, HargaPerBulan
}
