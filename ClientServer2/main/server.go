package main

import (
	"bufio"
	"fmt"
	"net"
)

// also can do multiple connections
// bc/ the server does not exit after the first run, can run client multiple times
func handleConnection(conn *net.Conn) {  // pointer to a network connection
	reader := bufio.NewReader(*conn)  // dereferences the pointer
	// any connection will wait in this loop -> good reason for that's asynchronous
	// won't hold any new connections to the server
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
		// server response to the client
		fmt.Fprintln(*conn, "OK")
	}

}


func main() {
	ln, _ := net.Listen("tcp", ":8030")
	for {
		conn, _ := ln.Accept()
		// we're on handleConnection, won't prevent this from continuing
		go handleConnection(&conn)
	}
}

// the server won't do anything until the client is run to send out
// the messages
//func main() {
//	ln, _ := net.Listen("tcp", ":8030")
//	conn, _ := ln.Accept()
//	reader := bufio.NewReader(conn)
//	msg, _ := reader.ReadString('\n')
//	fmt.Println(msg)
//}


//func main() {
//	msgPtr := flag.String("msg", "Default message", "Message to send")
//	flag.Parse()
//	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
//	fmt.Println(conn, msgPtr)
//}
