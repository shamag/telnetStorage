package main

import "fmt"
import storage "memoryStorage/lib/storage"

func main()  {
	var localStorage storage.SyncStorage
	store := localStorage.Create()
	store.Set("key", "value")
	store.Set("key", "value2")

	var localStorage2 storage.AtomicStorage
	store2 := localStorage2.Create()
	store2.Set("key", "value")
	fmt.Printf("key1 = %s \n", store2.Get("key"))
	for i:=0; i<10 ;i++  {
		go func(i int) {
			if i == 5 {
				store2.Set("key", "val4")
			}

		    fmt.Printf("key4 = %s  %d \n", store2.Get("key"), i)
		}(i)
	}
	fmt.Printf("key3 = %s \n", store2.Get("key"))


	fmt.Scanln()

}
