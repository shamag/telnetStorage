package main

import (
	"bufio"
	"fmt"
	storage "memoryStorage/lib/storage"
	"net"
)

func createHandler(store *storage.SyncStorage) func(net.Conn, string) {
	return func(conn net.Conn, com string) {
		conn.Write([]byte(com))
	}
}
func main() {
	var localStorage storage.SyncStorage
	store := localStorage.Create()
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}
	handler := createHandler(store)
	for {
		fmt.Printf("Listening at %q. \n", listener.Addr())
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		fmt.Printf("Received new connection from %q. \n", conn.RemoteAddr())
		command, _ := bufio.NewReader(conn).ReadString('\n')
		go handler(conn, command)
		fmt.Printf("Spawned handler to handle connection from %q. \n", conn.RemoteAddr())
	}
	// store.Set("key", "value")
	// store.Set("key", "value2")

	// var localStorage2 storage.AtomicStorage
	// store2 := localStorage2.Create()
	// store2.Set("key", "value")
	// fmt.Printf("key1 = %s \n", store2.Get("key"))
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		if i == 5 {
	// 			store2.Set("key", "val4")
	// 		}

	// 		fmt.Printf("key4 = %s  %d \n", store2.Get("key"), i)
	// 	}(i)
	// }
	// fmt.Printf("key3 = %s \n", store2.Get("key"))

	// fmt.Scanln()

}
