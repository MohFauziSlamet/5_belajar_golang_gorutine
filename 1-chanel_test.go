package belajar_golang_gorutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
//* Pengenalan Channel :
● Channel adalah tempat komunikasi secara synchronous
	yang bisa dilakukan oleh goroutine.
● Di Channel terdapat pengirim dan penerima, biasanya pengirim dan
	penerima adalah goroutine yang berbeda.
● Saat melakukan pengiriman data ke Channel, goroutine akan ter-block,
	sampai ada yang menerima data tersebut.
● Maka dari itu, channel disebut sebagai alat komunikasi synchronous
	(blocking).
● Channel cocok sekali sebagai alternatif seperti mekanisme async await
	yang terdapat di beberapa bahasa pemrograman lain.


//*	Karakteristik Channel :
● Secara default channel hanya bisa menampung satu data, jika kita ingin
	menambahkan data lagi, harus menunggu data yang ada di channel diambil.
● Channel hanya bisa menerima satu jenis data.
● Channel bisa diambil dari lebih dari satu goroutine.
● Channel harus di close jika tidak digunakan, atau bisa menyebabkan
	memory leak

//*	Membuat Channel :
● Channel di Go-Lang direpresentasikan dengan tipe data chan.
● Untuk membuat channel sangat mudah, kita bisa menggunakan make(),
	mirip ketika membuat map.
● Namun saat pembuatan channel, kita harus tentukan tipe data apa yang
	bisa dimasukkan kedalam channel tersebut


	channel := make(chan, string)


//*  Mengirim dan Menerima Data dari Channel
● Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk
	mengirim dan menerima data.
● Untuk mengirim data, kita bisa gunakan kode : channel <- data
● Sedangkan untuk menerima data, bisa gunakan kode : data <- channel
● Jika selesai, jangan lupa untuk menutup channel menggunakan
	function close().

*/

func TestCreateChanel(t *testing.T) {
	//* make chanel
	channel := make(chan string)

	//* methode defer , dijalankan ketika func selesai atau error.
	//* menutup chanel agar tidak terjadi memory leak.
	defer close(channel)

	//* make anonimous func
	go func() {
		time.Sleep(2 * time.Second)

		//* passing data to channel
		channel <- "Muhammad AzkaFitra Ramadhan"

		fmt.Println("Selesai mengirim data ke channel")

	}()

	//* receive data from channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	/*
		* Noted :
		ketika membuat channel , wajib pastikan ada kode untuk mengirim dan
		menerima data dari channel .
		jika hanya mengirim saja , tanpa ada yang menerima ,
		akan terjadi error. dan sebaliknya.

	*/
}

//? ==================================================================

/*
//* Channel Sebagai Parameter :
● Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim
channel ke function lain via parameter.
● Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah
pass by value, artinya value akan diduplikasi lalu dikirim ke function
parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan
pointer (agar pass by reference).
● Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut.
*/

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	//* passing data to channel
	channel <- "Muhammad AzkaFitra Ramadhan"

	fmt.Println("Selesai mengirim data ke channel")
}
func TestChannelAsParameter(t *testing.T) {
	// * make chanel
	channel := make(chan string)

	// * methode defer , dijalankan ketika func selesai atau error.
	// * menutup chanel agar tidak terjadi memory leak.
	defer close(channel)

	//* call [GiveMeResponse]
	go GiveMeResponse(channel)

	// * receive data from channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

//? ==================================================================

/*
//* Channel In dan Out :
● Saat kita mengirim channel sebagai parameter, isi function tersebut
bisa mengirim dan menerima data dari channel tersebut/
● Kadang kita ingin memberi tahu terhadap function, misal bahwa channel
tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan
untuk menerima data.
● Hal ini bisa kita lakukan di parameter dengan cara menandai apakah
channel ini digunakan untuk in (mengirim data) atau out (menerima data).
*/

// * channel in
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	//* jika memaksa untuk menerima, maka akan terjadi error.
	//* invalid operation: cannot receive from send-only
	//* channel channel (variable of type chan<- string)
	// data := <-channel

	//* send data to channel
	channel <- "Muhammad AzkaFitra Ramadhan"
	fmt.Println("Selesai mengirim data ke channel")

}

// * channel out
func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)

	//* jika memaksa untuk mengirim, maka akan terjadi error.
	//* invalid operation: cannot send from receive-only
	//* channel channel (variable of type <-chan string)
	// channel <- "Muhammad AzkaFitra Ramadhan"

	//* receive data from channel
	data := <-channel
	fmt.Println(data)
	fmt.Println("Selesai menerima data ke channel")
	fmt.Println()

}

func TestChannelInOut(t *testing.T) {
	// * make chanel
	channel := make(chan string)

	// * methode defer , dijalankan ketika func selesai atau error.
	// * menutup chanel agar tidak terjadi memory leak.
	defer close(channel)

	//* call [OnlyIn] to send data to channel
	go OnlyIn(channel)

	//* call [OnlyOut] to receive data from channel
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

//? ==================================================================

/*
//* Buffered Channel
● Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu
	hanya bisa menerima 1 data.
● Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu
	sampai data ke-1 ada yang mengambil.
● Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima,
	dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan
	ikut lambat juga
● Untungnya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk
	menampung data antrian di Channel.



//*	Buffer Capacity
● Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
● Ini cocok sekali ketika memang goroutine yang menerima data lebih
lambat dari yang mengirim
● Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
● Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai
buffer ada yang kosong datanya.
*/

func TestBufferedChannel(t *testing.T) {
	// * make chanel
	// channel := make(chan string) // without buffered
	channel := make(chan string, 3) // with buffered , and lenght 3.

	// * methode defer , dijalankan ketika func selesai atau error.
	// * menutup chanel agar tidak terjadi memory leak.
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Azkafitra"
		channel <- "Ramadhan"
	}()

	go func() {
		fmt.Printf("(<-channel): %v\n", (<-channel))
		fmt.Printf("(<-channel): %v\n", (<-channel))
		fmt.Printf("(<-channel): %v\n", (<-channel))
	}()

	/*
		Noted :
		Jika menggunakan buffer , walau data chanel belum ada yang mengambil
		tidak akan terjadi error . selagi buffer masih dapat menampung channel.

		jika buffer 3 , namun di isi dengan channel 1 . dan data belum di ambil,
		tidak akan terjadi error.
		namun jika buffer 3 , dan channel yang masuk ada 4 , maka akan error.
	*/

	fmt.Printf("cap(channel): %v\n", cap(channel))
	fmt.Printf("len(channel): %v\n", len(channel))

	time.Sleep(2 * time.Second)
}

//? ==================================================================

/*
Range Channel
● Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus
oleh pengirim
● Dan kadang tidak jelas kapan channel tersebut akan berhenti
menerima data
● Salah satu yang bisa kita lakukan adalah dengan menggunakan
perulangan range ketika menerima data dari channel
● Ketika sebuah channel di close(), maka secara otomatis perulangan
tersebut akan berhenti
● Ini lebih sederhana dari pada kita melakukan pengecekan channel
secara manual
*/

func TestRangeChannel(t *testing.T) {
	// * make chanel
	channel := make(chan string) // without buffered

	// * methode defer , dijalankan ketika func selesai atau error.
	// * menutup chanel agar tidak terjadi memory leak.
	defer close(channel)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
	}()

	go func() {
		for v := range channel {

			fmt.Printf("(<-channel): %v\n", (v))
		}
	}()

	/*
		Noted :
		Jika menggunakan buffer , walau data chanel belum ada yang mengambil
		tidak akan terjadi error . selagi buffer masih dapat menampung channel.

		jika buffer 3 , namun di isi dengan channel 1 . dan data belum di ambil,
		tidak akan terjadi error.
		namun jika buffer 3 , dan channel yang masuk ada 4 , maka akan error.
	*/

	time.Sleep(2 * time.Second)
}

//? ==================================================================

/*
Select Channel
● Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
● Lalu kita ingin mendapatkan data dari semua channel tersebut
● Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
● Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang
secara bersamaan di beberapa channel, maka akan dipilih secara random
*/

func TestSelectChannel(t *testing.T) {
	// * make chanel
	channel1 := make(chan string) // without buffered
	channel2 := make(chan string) // without buffered

	// * methode defer , dijalankan ketika func selesai atau error.
	// * menutup chanel agar tidak terjadi memory leak.
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	//* menjalankan secara manual
	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	//* menjalankan dengan for
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		//* kondisi untuk menghentikan perulangan
		if counter == 1 {
			break
		}

	}

}

//? ==================================================================

/*
Default Select
● Apa yang terjadi jika kita melakukan select terhadap channel
yang ternyata tidak ada datanya?.
● Maka kita akan menunggu sampai data ada.
● Kadang mungkin kita ingin melakukan sesuatu jika misal semua
channel tidak ada datanya ketika kita melakukan select channel.
● Dalam select, kita bisa menambahkan default, dimana ini akan
dieksekusi jika memang di semua channel yang kita select
tidak ada datanya
*/

func TestDefaultSelectChannel(t *testing.T) {
	// * make chanel
	channel1 := make(chan string) // without buffered
	channel2 := make(chan string) // without buffered

	// * methode defer , dijalankan ketika func selesai atau error.
	// * menutup chanel agar tidak terjadi memory leak.
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	//* menjalankan secara manual
	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data dari channel 2", data)
	// }

	//* menjalankan dengan for
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data . . . ")

		}

		//* kondisi untuk menghentikan perulangan
		if counter == 1 {
			break
		}

	}

}
