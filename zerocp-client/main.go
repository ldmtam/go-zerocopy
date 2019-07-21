package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
	"time"
)

var filePtr = flag.String("file", "", "file path")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()

	fileFd, err := syscall.Open(*filePtr, syscall.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fileFd)

	stat := new(syscall.Stat_t)

	if err = syscall.Fstat(fileFd, stat); err != nil {
		panic(err)
	}

	netFd, err := conn.(*net.TCPConn).File()
	if err != nil {
		panic(err)
	}
	defer netFd.Close()

	offset := int64(0)
	n, err := syscall.Sendfile(int(netFd.Fd()), fileFd, &offset, int(stat.Size))
	if err != nil {
		panic(err)
	}

	fmt.Printf("sent %v byte(s) in %v\n", n, time.Since(start).String())
}
