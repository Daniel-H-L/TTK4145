package main

import (
	"./Network"
	"fmt"
	"time"
)

func test_send(destinationIP string, msg chan string) {
	go Network.Udp_interface_send(destinationIP, msg)
	msg <- "Hola amiga"
	fmt.Println("Message sent")

}

func main() {
	ch := make(chan string)
	test_send("255.255.255.255", ch)
	time.Sleep(200 * time.Millisecond)
}
