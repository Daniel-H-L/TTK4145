package main

import (
	"./Network"
	"fmt"
	"time"
)

func test_broadcast(msg chan string) {
	go Network.Udp_interface_bcast(msg)

	for {
		msg <- "Hello!"
		time.Sleep(200 * time.Millisecond)
	}
}

// func test_send(destinationIP string, msg chan string) {
// 	msg <- "Hola amiga"
// 	go Network.Udp_interface_send(destinationIP, msg)
// }

func test_receive(msg chan string) {
	go Network.Udp_interface_receive(msg)

	time.Sleep(200 * time.Millisecond)
}

func main() {
	//ch := make(chan string)
	mg := make(chan string)
	//test_broadcast(ch)
	//test_send("255.255.255.255", mg)
	test_receive(mg)

	val := <-mg
	fmt.Println("Message received: " + val)

	time.Sleep(800 * time.Millisecond)
}
