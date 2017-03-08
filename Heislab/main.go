package main

import (
	"./DriveElevator"
	"./Master"
	"./Network"
	"./Slave"
	"fmt"
	"time"
)

var backup = Network.Backup{}
var MasterIP = ""

const BCAST_PORT = ":40018"
const PORT = ":30018"

func main() {
	chan_slave_received_msg := make(chan []byte, 1)
	chan_master_receive_msg := make(chan []byte, 1)

	chan_is_alive := make(chan string, 1)
	chan_master_bcast := make(chan string, 1)

	chan_error_slave := make(chan error, 1)
	chan_error_master := make(chan error, 1)

	chan_change_to_master := make(chan bool, 1)

	chan_descendant_nr := make(chan int, 1)

	chan_new_order := make(chan Network.NewOrder, 1)
	chan_elev_order := make(chan Driveelevator.Orders, 1)
	chan_order_executed := make(chan int, 1)

	Slave.Slave_init(chan_slave_received_msg, chan_master_bcast, BCAST_PORT, chan_error_slave, chan_change_to_master, &MasterIP)
	DriveElevator.Run_elevator()

	for {
		select {
		case change_state := <-chan_change_to_master:
			if change_state == 1 {
				//Må først gå i en sikker state. Dvs utfør alle interne ordre først. Ikke at i mot noen nye.
				Master.Master_detect_slaves(chan_master_receive_msg, chan_is_alive, PORT, chan_error_master)
			} else {
				Slave.Slave_drive_elevator(chan_new_order, chan_elev_order, chan_slave_received_msg, PORT, chan_error_slave, chan_change_to_master, chan_descendant_nr, chan_order_executed, &MasterIP)
			}
		case Network.Udp_get_local_ip() == MasterIP:
			Master.Master_drive_elevator()

		}
	}
}
