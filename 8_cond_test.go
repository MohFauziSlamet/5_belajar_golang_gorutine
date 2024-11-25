package belajar_golang_gorutine

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

/*
Cond
● Cond adalah adalah implementasi locking berbasis kondisi.
● Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex)
untuk implementasi locking nya, namun berbeda dengan Locker
biasanya, di Cond terdapat function Wait() untuk bertanya apakah
perlu menunggu atau tidak.
● Function Signal() bisa digunakan untuk memberi tahu sebuah
goroutine agar tidak perlu menunggu lagi, sedangkan function
Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak
perlu menunggu lagi
● Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
*/

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()

	log.Println("WaitGroup dijalankan", value)
	cond.L.Unlock()
}

func TestCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	fmt.Println("check flow 2")
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			// cond.Signal()
			cond.Broadcast()
		}
	}()

	group.Wait()
}
