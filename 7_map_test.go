package belajar_golang_gorutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
Map
● Go-Lang memiliki sebuah struct beranama sync.Map.
● Map ini mirip Go-Lang map, namun yang membedakan, Map ini
aman untuk menggunaan concurrent menggunakan goroutine.
● Ada beberapa function yang bisa kita gunakan di Map :
○ Store(key, value) untuk menyimpan data ke Map.
○ Load(key) untuk mengambil data dari Map menggunakan key.
○ Delete(key) untuk menghapus data di Map menggunakan key.
○ Range(function(key, value)) digunakan untuk melakukan iterasi
seluruh data di Map.
*/

func AddToMap(data *sync.Map, index int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	data.Store(index, index)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			AddToMap(data, i, group)
		}()
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, " : ", value)
		return true
	})
}
