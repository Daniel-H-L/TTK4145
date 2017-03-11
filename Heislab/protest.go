package main

import (
	"./src/DriveElevator"
	"./src/Master"
	"./src/Network"
	"./src/Slave"
	"fmt"
	"os/exec"
	"time"
)

var MasterIP = "129.241.187.141"

var local_ip = ""

const BCAST_PORT = ":40018"
const PORT = ":30018"

func main() {
	fmt.Println("Start main...")
	chan_state := make(chan int, 1)
	chan_dir := make(chan int, 1)
	chan_floor := make(chan int, 1)

	//chan_state_slave := make(chan Network.ElevState, 1)
	chan_order_executed := make(chan int, 10)

	chan_new_master_order := make(chan DriveElevator.Button, 1)
	chan_new_hw_order := make(chan DriveElevator.Button, 1)

	chan_source_ip := make(chan string, 1)

	chan_error_slave := make(chan error, 1)

	chan_set_lights := make(chan [3][4]int, 1)

	chan_is_alive := make(chan string, 1)
	chan_master_ip := make(chan string, 1)
	//chan_local_ip := make(chan string, 1)

	chan_master := make(chan Master.Master, 1)

	chan_new_network_order := make(chan Network.NewOrder, 1)
	chan_slave_received_msg := make(chan []byte, 1)
	chan_master_received_msg := make(chan []byte, 1)

	chan_change_to_master := make(chan bool, 1)
	chan_change_to_slave := make(chan bool, 1)

	chan_master_bcast := make(chan string, 1)
	chan_network_lights := make(chan [3][4]int, 1)

	chan_kill := make(chan bool)

	//chan_backup := make(chan Network.Backup, 1)
	chan_reset := make(chan bool, 1)
	chan_descendant_nr := make(chan int, 1)

	chan_slavelist := make(chan map[string]chan bool)

	chan_elev_state := make(chan Network.ElevState, 1)
	state := 2
	for {
		switch state {
		case 0:
		case 1:
			time.Sleep(5 * time.Second)
			Slave.Slave_init(chan_slave_received_msg, chan_master_bcast, BCAST_PORT, chan_error_slave, chan_change_to_master, &MasterIP, state, chan_new_network_order, chan_new_hw_order, chan_new_master_order, chan_order_executed, chan_set_lights, chan_network_lights, PORT, chan_source_ip, chan_descendant_nr, chan_new_network_order, chan_elev_state, chan_kill, chan_is_alive)
			fmt.Println("Slave init done...")
		case 2:
			Master.Master_init(chan_master, chan_master_received_msg, chan_is_alive, PORT, chan_change_to_slave, chan_master_ip, chan_new_hw_order, chan_new_master_order, chan_source_ip, chan_set_lights, chan_order_executed, chan_state, chan_dir, chan_floor, chan_elev_state, chan_reset, chan_descendant_nr, chan_slavelist, chan_network_lights, chan_new_network_order)

			cmd := exec.Command("gnome-terminal", "-x", "go", "run", "protest.go")
			cmd.Run()
			select {
			case <-time.After(20 * time.Second):
				state = 1
			}
		}
	}
}
