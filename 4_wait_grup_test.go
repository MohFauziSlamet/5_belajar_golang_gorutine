package belajar_golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
WaitGroup
● WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
● Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine,
tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
● Kasus seperti ini bisa menggunakan WaitGroup
● Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah
proses goroutine selesai, kita bisa gunakan method Done()
● Untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()
*/

func RunAsyncronous(group *sync.WaitGroup) {
	//? dijalankan ketika func selesai
	defer group.Done()

	//? eksekusi program
	group.Add(1)
	fmt.Println("Hello Azka")
	time.Sleep(2 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	//? set jadi pointer
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsyncronous(group)
	}

	group.Wait()
	fmt.Println("Complete")
}
