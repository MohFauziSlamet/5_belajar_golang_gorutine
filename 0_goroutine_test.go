package belajar_golang_gorutine

import (
	"fmt"
	"testing"
	"time"
)

/*
Membuat Goroutine
● Untuk membuat goroutine di Golang sangatlah sederhana
● Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan
dalam goroutine
● Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara
asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
● Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita
buat selesai

*/

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {

	//* ini adalah contoh menjalankan program
	//* secara squence , dari atas kebawah.
	// RunHelloWorld()
	// fmt.Println("Ups")

	//* ini adalah contoh menjalankan dengan goroutine
	go RunHelloWorld()
	fmt.Println("Ups")

	/*
		ketika menjalankan dengan menggunakan goroutine ,
		maka yang dijalankan tsb akan menjadi asyncrounus (tidak akan ditunggu
		sampai func atau kode itu selesai dijalankan).

		Berhati-hati lah menggunakan goroutine pada func yang mengembalikan
		value. karena ada kemunkinan value return tsb tidak akan dapat , karena
		func atau kode terlanjur di kill sama go-scheduler , dikarenakan lama.
	*/

	//* coba gunakan delay untuk menunggu func atau kode itu dijalankan,
	//* namun jika sudah di berikan delay , namun masih lama , maka akan
	//* tetap di kill oleh go-scheduler.
	//*
	//* menambah delay dengan package time.
	time.Sleep(500 * time.Millisecond)
}

//* ============================================================

/*
Goroutine Sangat Ringan
● Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
● Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
● Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan
*/

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {

	//* kita loop sampai 100.000 kali
	for i := 0; i < 100000; i++ {
		// //* menjalankan func tanpa goroutine
		// DisplayNumber(i) // hasilnya akan berurutan

		//* menjalankan func dengan goroutine
		go DisplayNumber(i) // hasilnya tidak akan berurutan
	}
}
