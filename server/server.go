// bind port
// listen to port
// accept conncetion
// read socket
// write socket

package main

import (
	. "fmt"
	"net"
	"os"
)

const (
	SERVER_HOST     = "localhost"
	SERVER_PORT     = "8080"
	SERVER_PROTOCOL = "tcp"
)

func main() {
	Println("server is running")
	ser, err := net.Listen(SERVER_PROTOCOL, Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	if err != nil {
		panic(err)
	}
	defer ser.Close()
	Println("Im listining to ", SERVER_PORT)

	for {
		conn, err := ser.Accept()
		if err != nil {
			Println("Conncetion with client failed")
			os.Exit(1)
		}
		Println("Client Connected")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	buff := make([]byte, 1024)
	rn, err := conn.Read(buff)
	if err != nil {
		Println("Cant Read Data from Client")
	}
	Println("Read data from client is: ", string(buff[:rn]))

	_, err = conn.Write([]byte("hello client, how u been?"))
	if err != nil {
		Println("Error writing data in socket: ", err.Error())
	}
	conn.Close()
}
