package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second)) // will close after 10 seconds

	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		io.WriteString(conn, "I heard you say "+ln)
	}
	defer conn.Close()

	fmt.Println("*******Got here!!!********")
}
