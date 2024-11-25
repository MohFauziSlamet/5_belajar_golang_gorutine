package belajar_golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Mutex (Mutual Exclusion)
● Untuk mengatasi masalah race condition tersebut, di Go-Lang
terdapat sebuah struct bernama sync.Mutex
● Mutex bisa digunakan untuk melakukan locking dan unlocking,
dimana ketika kita melakukan locking terhadap mutex, maka
tidak ada yang bisa melakukan locking lagi sampai kita
melakukan unlock
● Dengan demikian, jika ada beberapa goroutine melakukan lock
terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan,
setelah goroutine tersebut melakukan unlock, baru goroutine
selanjutnya diperbolehkan melakukan lock lagi
● Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
*/
func TestMutex(t *testing.T) {

	var mutex sync.Mutex
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()

			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("counter x = ", x)

}

/*
RWMutex (Read Write Mutex)
● Kadang ada kasus dimana kita ingin melakukan locking tidak
hanya pada proses mengubah data, tapi juga membaca data.
● Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya
nanti akan rebutan antara proses membaca dan mengubah.
● Di Go-Lang telah disediakan struct RWMutex (Read Write Mutex)
untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock,
lock untuk Read dan lock untuk Write.
*/

type BankAccount struct {
	// RWMutex sync.RWMutex //penulisan boleh di beri init variable
	sync.RWMutex // boleh langsung type nya , jika itu struct
	Balance      int
}

/*
(account *BankAccount) , ini menunjukan suatu penerimna dari func
addBalance. dengan kata lain , func addBalance di tambahkan ke dalam
struct BankAccount.
*/
func (account *BankAccount) addBalance(amount int) {
	account.RWMutex.Lock() // lock untuk write
	defer account.RWMutex.Unlock()

	account.Balance = account.Balance + amount
}

func (account *BankAccount) getBalance() int {
	account.RWMutex.RLock() // lock untuk read
	defer account.RWMutex.RUnlock()

	balance := account.Balance
	return balance

}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 1000; i++ {
		go func() {

			for j := 1; j <= 100; j++ {
				account.addBalance(1)
				fmt.Printf("account.getBalance(): %v\n", account.getBalance())
			}
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("counter x = ", account.getBalance())

	/*
		Jika struct tidak pernah dipanggil dengan menggunakan
		goroutine , maka tidak perlu menambahkan syncronisasi (mutex).
	*/
}

//? ================================================

/*
* Deadlock
● Hati-hati saat membuat aplikasi yang parallel atau concurrent,
masalah yang sering kita hadapi adalah Deadlock.
● Deadlock adalah keadaan dimana sebuah proses goroutine saling
menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
● Sekarang kita coba simulasikan proses deadlock.
*/

type UserBalance struct {
	sync.RWMutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.RWMutex.Lock()
}

func (user *UserBalance) UnLock() {
	user.RWMutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(FromUser *UserBalance, ToUser *UserBalance, amount int) {
	defer FromUser.UnLock()
	defer ToUser.UnLock()

	FromUser.Lock()
	fmt.Println("Lock FromUser: ", FromUser.Name)
	FromUser.Change(-amount)

	time.Sleep(1 * time.Second)

	ToUser.Lock()
	fmt.Println("Lock ToUser: ", ToUser.Name)
	ToUser.Change(amount)

	time.Sleep(1 * time.Second)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Azka",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Fitra",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 100000)

	time.Sleep(2 * time.Second)

	fmt.Println("User 1 : ", user1.Name, "Balance: ", user1.Balance)
	fmt.Println("User 2 : ", user2.Name, "Balance: ", user2.Balance)
}
