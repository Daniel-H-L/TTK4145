package main

import (
	"./Network"
	"fmt"
	"os/exec"
	"time"
)

var State int
var localIP string
var queue = make([]int, 2)

func start() {
	buffer := make([]byte, 1024)
	conn := Network.Udp_interface_init(":40018")

	conn.SetReadDeadline(time.Now().Add(time.Second))
	_, _, err := conn.ReadFromUDP(buffer)

	if err != nil {
		State = 1
		fmt.Println("I am master...")
		cmd := exec.Command("gnome-terminal", "-x", "go", "run", "testpair.go")
		cmd.Run()
	} else {
		State = 2
		fmt.Println("I am slave...")
	}
	conn.Close()
}

func master_bcast(message Network.StandardData) {
	Network.Udp_broadcast("", message.Order_executed)
	time.Sleep(200 * time.Millisecond)
}

func slave_backup(ex_chan chan int, rec_chan chan []byte, err_chan chan error, state int) {
	go Network.Udp_receive_order_executed(ex_chan, rec_chan, ":40018", err_chan, state)
	fmt.Println("Waiting for error\n")
	err := <-err_chan
	fmt.Println("Error received\n")
	if err != nil {
		State = 0
	}
}

func main() {
	State = 0

	rec_chan := make(chan []byte, 1)
	exec_chan := make(chan int, 1)
	err_chan := make(chan error, 1)

	message := Network.StandardData{}
	message.Order_executed = 3

	for {
		switch State {
		case 0:
			fmt.Println("Start\n")

			start()

		case 1:
			for {
				time.Sleep(200 * time.Millisecond)
				fmt.Println("Master ready to send...")
				master_bcast(message)
				fmt.Println("Sent: ", message.Order_executed)
			}
		case 2:
			fmt.Println("Slave\n")
			slave_backup(exec_chan, rec_chan, err_chan, State)

		}
	}
}
