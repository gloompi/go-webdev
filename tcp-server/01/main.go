package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(con, "\n Hello from TCP server")
		con.Close()
	}
}
