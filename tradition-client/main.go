package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
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

	b, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		panic(err)
	}

	n, err := conn.Write(b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("sent %v byte(s) in %s\n", n, time.Since(start).String())
}
