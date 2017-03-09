package main

import (
	"./src/DriveElevator"
	"./src/Master"
	"./src/Network"
	//"./src/Slave"
	//"fmt"
	"time"
)

var backup = Network.Backup{}

//var master = Master.Master{}
var MasterIP = "129.241.187.142"

//var local_ip = ""

const BCAST_PORT = ":40018"
const PORT = ":30018"

func main() {

	chan_state := make(chan int, 1)
	chan_dir := make(chan int, 1)
	chan_floor := make(chan int, 1)

	chan_state_slave := make(chan Network.ElevState, 1)
	chan_order_executed := make(chan int, 1)
	chan_kill := make(chan bool)
	//chan_new_network_order := make(chan Network.NewOrder, 1)
	chan_new_master_order := make(chan DriveElevator.Button, 1)
	chan_new_hw_order := make(chan DriveElevator.Button, 1)

	chan_source_ip := make(chan string, 1)
	chan_master_receive_msg := make(chan []byte, 1)

	//chan_error_master := make(chan error, 1)

	chan_set_lights := make(chan [][]int, 1)

	go DriveElevator.Run_elevator(chan_state, chan_dir, chan_floor, chan_order_executed, chan_new_hw_order, chan_new_master_order, chan_set_lights)

	go Master.Master_test_drive(chan_new_hw_order, chan_new_master_order, &backup, chan_source_ip, chan_set_lights, chan_order_executed, chan_state, chan_dir, chan_floor, chan_kill, &MasterIP, chan_master_receive_msg, chan_state_slave)

	for {
		time.Sleep(50 * time.Millisecond)
	}

	/*fmt.Println("Start main...")
	chan_slave_received_msg := make(chan []byte, 1)
	chan_master_receive_msg := make(chan []byte, 1)

	chan_source_ip := make(chan string, 1)
	chan_master_ip := make(chan string, 1)

	chan_is_alive := make(chan string, 1)
	chan_master_bcast := make(chan string, 1)

	chan_error_slave := make(chan error, 1)
	chan_error_master := make(chan error, 1)

	chan_change_to_master := make(chan bool, 1)

	chan_descendant_nr := make(chan int, 1)

	chan_new_order := make(chan Network.NewOrder, 1)
	chan_new_hw_order := make(chan EventManager.Orders)
	//chan_elev_order := make(chan EventManager.Orders, 1)
	chan_order_executed := make(chan int, 1)

	chan_state := make(chan int, 1)
	chan_dir := make(chan int, 1)
	chan_floor := make(chan int, 1)

	chan_new_network_order := make(chan EventManager.Orders,1)
	chan_new_hw_order := make(chan EventManager.Orders, 1)

	fmt.Println("In main...")

	// Slave.Slave_init(chan_slave_received_msg, chan_master_bcast, BCAST_PORT, chan_error_slave, chan_change_to_master, &MasterIP, state int)
	// fmt.Println("Slave init done...")

	go DriveElevator.Run_elevator(chan_state, chan_dir, chan_floor, chan_order_executed, chan_new_hw_order, chan_new_network_order)

	//local_ip,_ := Network.Udp_get_local_ip()
	//chan_master_ip <- MasterIP
	state := 0
	for {
		switch state{
		case 0:
			Slave.Slave_init(chan_slave_received_msg, chan_master_bcast, BCAST_PORT, chan_error_slave, chan_change_to_master, &MasterIP, state)
			fmt.Println("Slave init done...")
			change_state := <-chan_change_to_master

			if change_state == true {
				time.Sleep(10*time.Millisecond)
				chan_kill <- 1
				//Må først gå i en sikker state. Dvs utfør alle interne ordre først. Ikke at i mot noen nye.
				state = 1
			} else {
				Slave.Slave_drive_elevator(chan_new_order, chan_new_hw_order, chan_slave_received_msg, PORT, chan_error_slave, chan_change_to_master, chan_descendant_nr, chan_order_executed, &MasterIP, chan_source_ip, chan_new_hw_order, chan_new_network_order)
			}
		case 1:
			Master.Master_detect_slave(chan_master_receive_msg, chan_is_alive, PORT, chan_error_master, &master)
			if local_ip == MasterIP {
				Master.Master_drive_elevator(&backup, chan_new_order, chan_new_hw_order, chan_source_ip, chan_master_receive_msg, PORT, chan_error_master, &master, chan_order_executed, chan_new_hw_order, chan_new_network_order)
			}

		}
	}*/
}
