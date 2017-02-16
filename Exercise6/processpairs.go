package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

var state int
var localIP string
var portNr string = ":30018"
var count int = -1
var queue = make([]int, 1)

type MSG struct {
	Msg string
	Cnt int
}

func Udp_get_local_ip() string {
	if localIP == "" {
		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{8, 8, 8, 8}, Port: 31800})
		if err != nil {
			return ""
		}
		defer conn.Close()
		localIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	fmt.Println("local ip speaking... ")
	return localIP
}

func Udp_interface_check_error(err error) {
	if err != nil {
		fmt.Println("... Error: ", err)
	}
}

func Udp_interface_init() *net.UDPConn {
	//socket
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	Udp_interface_check_error(err)

	conn, err := net.ListenUDP("udp", localAddr)
	Udp_interface_check_error(err)

	return conn
}

func start() {
	buffer := make([]byte, 1024)
	conn := Udp_interface_init()

	conn.SetReadDeadline(time.Now().Add(time.Second))
	_, _, err := conn.ReadFromUDP(buffer)

	if err != nil {
		state = 1
		fmt.Println("I am master...")
		cmd := exec.Command("gnome-terminal", "-x", "go", "run", "processpairs.go")
		cmd.Run()
	} else {
		state = 2
		fmt.Println("I am slave...")
	}
	conn.Close()
}

func primary_bcast(message MSG) {
	fmt.Println("Master count: ", message.Cnt)
	localAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:31800")
	Udp_interface_check_error(err)
	conn, err := net.DialUDP("udp", nil, localAddr)
	Udp_interface_check_error(err)
	defer conn.Close()
	json_ob, _ := json.Marshal(message)

	if len(json_ob) > 0 {
		conn.Write(json_ob)
	}
	time.Sleep(200 * time.Millisecond)

}

func backup() {
	buffer := make([]byte, 1024)
	addr, _ := net.ResolveUDPAddr("udp", ":31800")
	socket, _ := net.ListenUDP("udp", addr)
	socket.SetReadDeadline(time.Now().Add(time.Second))

	mlen, _, err := socket.ReadFromUDP(buffer)
	fmt.Println("Received: ", string(buffer[0:]))

	if err != nil {
		count = queue[1]
		state = 0
	} else {
		queue := queue[1:]
		rec_msg := MSG{}
		json.Unmarshal(buffer[:mlen], &rec_msg)
		i := rec_msg.Cnt
		fmt.Println("Backup received ", i)
		queue := append(queue, i)
	}
	socket.Close()
}

func main() {
	state = 0

	message := MSG{Udp_get_local_ip(), count}

	for {
		switch state {
		case 0:
			start()
		case 1:
			for {
				count++
				message.Cnt = count
				time.Sleep(200 * time.Millisecond)
				primary_bcast(message)
				fmt.Println("Count: ", count)
			}
		case 2:
			backup()
		}

	}

}
