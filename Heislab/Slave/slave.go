package Slave

import (
	"./DriveElevator"
	"./Network"
	"time"
)

var Master_IP string
var Descendant_nr int
var Change_to_master bool

func Slave_init(chan_received_msg chan []byte, chan_is_alive chan string, portNr string, chan_error chan error) {
	Master_IP = slave_listen_for_master(chan_received_msg, chan_is_alive, portNr, chan_error)
}

func Slave_send_is_alive() {
	Network.Udp_send_is_alive(Master_IP)
}

func slave_listen_for_master(chan_received_msg chan []byte, chan_is_alive chan string, portNr string, chan_error chan error, chan_change_to_master chan bool) string {
	go Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr, chan_error)
	ip := <-chan_is_alive

	if ip == string() {
		if Descendant_nr == 1 {
			1 <- chan_change_to_master
		} else if Descendant_nr == 2 {
			time.Sleep(5 * time.Second)
			Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr, chan_error)
			ip := <-chan_is_alive
			if ip == string() {
				1 <- chan_change_to_master
			}
		}
	} else {
		0 <- chan_change_to_master
	}
	return ip
}

func Slave_drive_elevator(chan_new_order chan Network.NewOrder, chan_elev_order chan Driveelevator.Orders, chan_received_msg chan []byte, port_nr string, chan_error chan error, chan_change_to_master chan bool) {
	go slave_receive_order_from_hw(chan_elev_order)
	go Network.Udp_receive_new_order(chan_new_order, chan_received_msg, port_nr, chan_error)
	for {
		select {
		case order := <-chan_new_order:
			if order != nil {
				DriveElevator.Eventmanager_add_new_order(order.floor, order.direction)
			}
		case elev_order := <-chan_elev_order:
			if elev_order != nil {
				order := Network.NewOrder{elev_order.floor, elev_order.dir, elev_order.is_inside, 1, 0, 0}
				Slave_send_new_order(order)
			}
		case master_order := <-chan_new_order:
			if chan_order != nil {
				DriveElevator.Eventmanager_add_new_order(master_order.floor, master_order.direction)
			}
		case change_to_master := <-chan_change_to_master:
		}
	}
}

func Slave_send_order_executed(chan_order_executed chan int) {
	DriveElevator.Driveelevator_order_executed(chan_order_executed)
	floor := <-chan_order_status

	order := NewOrder{}
	order.floor = floor
	order.is_new = false
	order.is_executed = true
	order.in_progess = false
	Network.Udp_send_order_status(order, Master_IP)
}

func slave_receive_order_from_hw(chan_new_order chan Driveelevator.Orders) {
	for {
		DriveElevator.Driveelevator_get_new_order(chan_new_order)
		new_order := <-chan_new_order
	}
}

func Slave_send_new_order(order NewOrder) {
	Network.Udp_send_new_order(order, Master_IP)
}

func Slave_receive_descendant(chan_descendant_nr chan int, chan_received_msg chan []byte, port_nr string, chan_error chan error) {
	go Network.Udp_receive_descendant_nr(chan_descendant_nr, chan_received_msg, port_nr, chan_error)
	err := <-chan_error

	if err == nil {
		Descendant_nr := <-chan_descendant_nr
	}
}

func Slave_mechanical_problem_send_to_master(order NewOrder) {
	if Drive_elevator.Eventmanager_stop_cause_mechanical_problem() == true {
		order.in_progress = false
		order.is_executed = false
		Network.Udp_send_order_status(order, Master_IP)
	}
}

func Slave_run() bool {
	Slave_init(chan_received_msg, chan_is_alive, ":40018", chan_error)

	if Network.Udp_get_local_ip() > Master_IP {
		return true
	} else {
		Slave_send_is_alive()
		Slave_receive_descendant(chan_descendant_nr, chan_received_msg, port_nr, chan_error)
	}

}
