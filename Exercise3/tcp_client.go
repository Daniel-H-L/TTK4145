package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:20018")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, text+"\n")

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message for server: " + msg)
	}
}
