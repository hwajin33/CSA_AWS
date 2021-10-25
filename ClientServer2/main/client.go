package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn)  {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	//msgPtr := flag.String("msg", "Default messages", "message to send")
	//flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	for {
		fmt.Println("Enter text: ")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, text)
		read(&conn)
	}

}

//func main() {
//	msg := "Here is a message"
//	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
//	fmt.Fprintf(conn, msg)
//}

//func handleConnection(conn *net.Conn) {
//	reader := bufio.NewReader(*conn)
//	msg, _ := reader.ReadString('\n')
//	fmt.Println(msg)

//}

//func main() {
//	ln, _ := net.Listen("tcp", ":8030")
//	for {
//		conn, _ := ln.Accept()
//		go handleConnection(&conn)
//	}
//}


