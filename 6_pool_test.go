package belajar_golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
Pool
● Pool adalah implementasi design pattern bernama object pool
pattern.
● Sederhananya, design pattern Pool ini digunakan untuk menyimpan
data, selanjutnya untuk menggunakan datanya, kita bisa mengambil
dari Pool, dan setelah selesai menggunakan datanya,
kita bisa menyimpan kembali ke Pool nya. pool sering digunakan
untuk menyimpan koneksi ke database.
● Implementasi Pool di Go-Lang ini sudah aman dari problem
race condition.
*/
func TestPool(t *testing.T) {
	//? membuat default pool
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}

	group := &sync.WaitGroup{}

	// Convert strings to pointer types before putting them in the pool
	azka := "Azka"
	fitra := "Fitra"
	ramadhan := "Ramadhan"
	pool.Put(&azka)
	pool.Put(&fitra)
	pool.Put(&ramadhan)

	for i := 0; i < 10; i++ {
		group.Add(1)

		go func() {
			data := pool.Get()

			// Type assert the data back to a *string
			if str, ok := data.(*string); ok {
				fmt.Println(*str)

				// Put it back as a pointer
				pool.Put(str)
			}

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Complete")
}
