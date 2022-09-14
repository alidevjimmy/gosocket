// connect to socket connection
// write into socket
// read

package main

import (
	"fmt"
	. "fmt"
	"net"
)

const (
	SERVER_HOST     = "localhost"
	SERVER_PORT     = "8080"
	SERVER_PROTOCOL = "tcp"
)

func main() {
	conn, err := net.Dial(SERVER_PROTOCOL, Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err.Error())
		}
		fmt.Println("Connection with server closed")
	}()
	Println("Connection with server established")
	n, err := conn.Write([]byte("Hello Server"))
	Println("n: ", n)
	buff := make([]byte, 1024)

	rn, err := conn.Read(buff)
	if err != nil {
		Println("Error while reading data from socker: ", err.Error())
	}
	Println("Read Data from server is: ", string(buff[:rn]))
}
