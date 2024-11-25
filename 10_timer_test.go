package belajar_golang_gorutine

import (
	"log"
	"sync"
	"testing"
	"time"
)

/*
Timer
● Timer adalah representasi satu kejadian
● Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel.
● Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration).
*/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	log.Println(time.Now())

	time := <-timer.C

	log.Println(time)

}

/*
time.After()
● Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
● Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)
*/

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	log.Println(time.Now())

	time := <-channel
	log.Println(time)

}

/*
time.AfterFunc()
● Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu.
● Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
● Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan
dipanggil ketika Timer mengirim kejadiannya.
*/

func TestAfterFunc(t *testing.T) {
	var group sync.WaitGroup

	group.Add(1)
	log.Println(time.Now())

	/// ini berjalan secara goroutine (asyncronous)
	time.AfterFunc(5*time.Second, func() {
		log.Println(time.Now())
		group.Done()
	})

	group.Wait()
}
