package belajar_golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
Once
● Once adalah fitur di Go-Lang yang bisa kita gunakan untuk
memastikan bahsa sebuah function dieksekusi hanya sekali
● Jadi berapa banyak pun goroutine yang mengakses, bisa
dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi
function nya
● Goroutine yang lain akan di hiraukan, artinya function
tidak akan dieksekusi lagi
*/

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i <= 10; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	t.Log("counter", counter)
	fmt.Println(counter)
}
