//https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":20018")

	conn, _ := ln.Accept()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received; ", string(msg))
		new_msg := strings.ToUpper(msg)

		conn.Write([]byte(new_msg + "\n"))
	}
}
