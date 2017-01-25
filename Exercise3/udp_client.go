package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("... Error: ", err)
	}
}

func main() {
	//Create socket (server)
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:23000")
	CheckError(err)

	//Create socket (client)
	ClientAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12000")
	CheckError(err)

	//Create dialog (datagram?)
	Conn, err := net.DialUDP("udp", ClientAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()

	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		ClientBuf := []byte(msg)
		_, err := Conn.Write(ClientBuf)

		if err != nil {
			fmt.Println(msg, err)
		}

		time.Sleep(time.Second * 1)
	}
}
