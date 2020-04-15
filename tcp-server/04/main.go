package main

import (
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	io.WriteString(conn, "I dialed you")
}
