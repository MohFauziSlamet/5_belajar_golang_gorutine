package belajar_golang_gorutine

import (
	"fmt"
	"testing"
	"time"
)

/*
*Masalah Dengan Goroutine
● Saat kita menggunakan goroutine, dia tidak hanya berjalan
secara concurrent, tapi bisa parallel juga, karena bisa ada
beberapa thread yang berjalan secara parallel.
● Hal ini sangat berbahaya ketika kita melakukan manipulasi
data variable yang sama oleh beberapa goroutine secara bersamaan.
● Hal ini bisa menyebabkan masalah yang namanya Race Condition.
*/
func TestRaceCondition(t *testing.T) {

	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("counter x = ", x)

	/*
		nilai counter tidak akan sampai 100.000
		dikarenakan ada goroutine yang mengakse nilai yang sama
		secara bersamaan.
		contoh , :
		x = 5
		x = 5 + 1
		x = 5 + 1
		x = 5 + 1
		x = 5 + 1

		ada goroutine yang menjalankan x = 5 + 1 secara bersamaan,
		bisa puluhan bahkan ratusan goroutine yang sama.
	*/
}
