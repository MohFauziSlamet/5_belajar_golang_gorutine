package belajar_golang_gorutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Atomic
● Go-Lang memiliki package yang bernama sync/atomic.
● Atomic merupakan package yang digunakan untuk menggunakan data
primitive secara aman pada proses concurrent.
● Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan
locking ketika ingin menaikkan angka di counter. Hal ini sebenarnya
bisa digunakan menggunakan Atomic package.
● Ada banyak sekali function di atomic package, kita bisa eksplore
sendiri di halaman dokumentasinya.
● https://golang.org/pkg/sync/atomic/
*/

func TestAtomic(t *testing.T) {
	var group sync.WaitGroup

	var x int64 = 0

	for i := 1; i <= 1000; i++ {
		group.Add(1)

		go func() {
			for j := 1; j <= 100; j++ {
				/// aman dari race condition
				atomic.AddInt64(&x, 1)

				/// kena race condition
				// x = x + 1
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println("counter x = ", x)

}
