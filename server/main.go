package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("listening on port 3000")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 4096)
	received := 0
	for n, err := conn.Read(b); err == nil && n > 0; n, err = conn.Read(b) {
		received += n
	}
	fmt.Printf("received %v byte(s) from %s\n", received, conn.RemoteAddr().String())
}
