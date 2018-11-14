package main

import (
	"bufio"
	"fmt"
	storage "memoryStorage/lib/storage"
	"net"
	"strings"
)

type handlers map[string]func(store storage.IStorage, args ...string) string

func createHandler(store storage.IStorage, hndl handlers) func(net.Conn) {
	return func(conn net.Conn) {
		defer conn.Close()
		for {
			mesaage, _ := bufio.NewReader(conn).ReadString('\n')
			args := strings.Fields(mesaage)
			if len(args) < 1 {
				continue
			}
			command := args[0]
			if strings.ToLower(command) == "exit" || strings.ToLower(command) == "close" {
				break
			}
			if len(args) < 2 {
				continue
			}
			args = args[1:]
			if handler, exist := hndl[command]; exist {
				result := handler(store, args...)
				conn.Write([]byte(result + "\n"))
			}
		}
	}
}
func main() {
	var localStorage storage.AtomicStorage
	store := localStorage.Create()
	listener, err := net.Listen("tcp", ":9990")
	defer listener.Close()
	if err != nil {
		panic(err)
	}
	var hndl = make(handlers)
	hndl["set"] = func(store storage.IStorage, args ...string) string {
		if len(args) < 2 {
			return "Wrong arguments number"
		}
		store.Set(args[0], args[1])
		return "OK"
	}
	hndl["get"] = func(store storage.IStorage, args ...string) string {
		if len(args) < 1 {
			return "Wrong arguments number"
		}
		var result string
		val, exist := store.Get(args[0])
		if exist {
			result = fmt.Sprintf("1\n%s", val)
		} else {
			result = "0"
		}
		return result
	}
	hndl["exist"] = func(store storage.IStorage, args ...string) string {
		if len(args) < 1 {
			return "Wrong arguments number"
		}
		if store.Exist(args[0]) {
			return "1"
		}
		return "0"
	}
	hndl["del"] = func(store storage.IStorage, args ...string) string {
		if len(args) < 1 {
			return "Wrong arguments number"
		}
		if store.Delete(args[0]) {
			return "1"
		}
		return "0"
	}
	handler := createHandler(store, hndl)
	for {
		fmt.Printf("Listening at %q. \n", listener.Addr())
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		rAddress := conn.RemoteAddr()
		fmt.Printf("Received new connection from %q. \n", rAddress)
		go handler(conn)
		fmt.Printf("Spawned handler to handle connection from %q. \n", rAddress)
	}

}
