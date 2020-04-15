package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	url := request(conn)
	response(conn, url)
}

func request(conn net.Conn) string {
	i := 0
	var url string

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)

		if i == 0 {
			url := strings.Fields(ln)[1]
			return url
		}
		if ln == "" {
			break
		}
		i++
	}

	return url
}

func response(conn net.Conn, url string) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Yo! Wassap everybody</title>
	</head>
	<body>
	<h1>` + url + `</h1>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
