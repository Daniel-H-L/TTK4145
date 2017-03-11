// package main

// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// const (
// 	INIT   = 0
// 	MASTER = 1
// 	SLAVE  = 2
// )

// const BCAST_PORT = ":40018"
// const PORT = ":30018"

// var MasterIP = "129.241.187.146"

// var state int

// func listenForMaster() {
// 	buffer := make([]byte, 1024)
// 	localAddr, err := net.ResolveUDPAddr("udp", BCAST_PORT)

// 	conn, err := net.ListenUDP("udp", localAddr)

// 	conn.SetReadDeadline(time.Now().Add(time.Second))
// 	_, _, err = conn.ReadFromUDP(buffer)

// 	if err != nil {
// 		state = MASTER
// 		fmt.Println("I am master...")

// 	} else {
// 		state = SLAVE
// 		fmt.Println("I am slave...")
// 	}
// 	conn.Close()
// }

// func main() {
// 	state = INIT

// 	fmt.Println("Start main...")

// 	listenForMaster()

// 	switch state {
// 	case MASTER:

// 		bcastAddr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:40018")

// 		conn, _ := net.DialUDP("udp", nil, bcastAddr)
// 		defer conn.Close()

// 		for {
// 			conn.Write([]byte(MasterIP))
// 			time.Sleep(200 * time.Millisecond)
// 		}

// 	case SLAVE:

// 		localAddr, _ := net.ResolveUDPAddr("udp", BCAST_PORT)

// 		conn, _ := net.ListenUDP("udp", localAddr)

// 		buffer := make([]byte, 1024)

// 		fmt.Println("In received...")
// 		for {

// 			// select {
// 			// case <-chan_kill:
// 			// 	fmt.Println("Killing interface")
// 			// 	conn.Close()
// 			// 	return
// 			// }

// 			conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

// 			n, _, err := conn.ReadFromUDP(buffer)

// 			if err != nil {
// 				fmt.Println("Received msg...", buffer[0:n])

// 			}
// 			fmt.Println("Add buffer to chan...")
// 		}
// 		defer conn.Close()
// 	}

// }
