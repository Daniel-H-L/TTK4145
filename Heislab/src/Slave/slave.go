package Slave

import (
	"../DriveElevator"
	"../DriveElevator/EventManager"
	"../Network"
	"time"
	"fmt"
)

var Descendant_nr int 

func Slave_init(chan_received_msg chan []byte, chan_master_bcast chan string, portNr string, chan_error chan error, chan_change_to_master chan bool, Master_IP *string, state int) {
	*Master_IP = slave_listen_for_master(chan_received_msg, chan_master_bcast, portNr, chan_error, chan_change_to_master)
	fmt.Println("In slave init...")
	local_ip,_ := Network.Udp_get_local_ip()
	err := <- chan_error
	if err != nil {
		chan_change_to_master <- true
		state = 1
	} else {
		if local_ip >= *Master_IP {
		chan_change_to_master <- true
		*Master_IP,_ = Network.Udp_get_local_ip()
		fmt.Println("Change to master...")
	} else {
		for {
			Network.Udp_send_is_alive(*Master_IP)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Remain as slave...")
			}
		}
	}
}

func slave_listen_for_master(chan_received_msg chan []byte, chan_is_alive chan string, portNr string, chan_error chan error, chan_change_to_master chan bool) string {
	go Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr, chan_error)
	ip := <-chan_is_alive

	if ip == "" {
		if Descendant_nr == 1 {
			chan_change_to_master <- true
		} else if Descendant_nr == 2 {
			time.Sleep(5 * time.Second)
			Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr, chan_error)
			ip := <-chan_is_alive
			if ip == "" {
				chan_change_to_master <- true
			}
		}
	} else {
		chan_change_to_master <- false
	}
	return ip
}

func Slave_drive_elevator(chan_new_order chan Network.NewOrder, chan_elev_order chan EventManager.Orders, chan_received_msg chan []byte, port_nr string, chan_error chan error, chan_change_to_master chan bool, chan_descendant_nr chan int, chan_order_executed chan int, Master_IP *string, chan_source_ip chan string, chan_new_hw_order chan EventManager.Orders, chan_new_network_order chan EventManager.Orders) {
	go slave_receive_order_from_hw(chan_elev_order)
	go Network.Udp_receive_new_order(chan_new_order, chan_received_msg, port_nr, chan_error, chan_source_ip)
	go Network.Udp_receive_descendant_nr(chan_descendant_nr, chan_received_msg, port_nr, chan_error)
	go DriveElevator.Driveelevator_order_executed(chan_order_executed)

	for {
		select {
		case elev_order := <-chan_elev_order:
			if elev_order != (EventManager.Orders{}) {
				order := Network.NewOrder{elev_order.Floor, elev_order.Dir, 1, 0, 0}
				Network.Udp_send_new_order(order, *Master_IP)
			}
		case master_order := <-chan_new_order:
			if master_order != (Network.NewOrder{}) {
				EventManager.Eventmanager_add_new_order(master_order.Floor, master_order.Direction)
			}
		case change_to_master := <-chan_change_to_master:
			if change_to_master == true {
				return
			}
		case descendant := <- chan_descendant_nr:
			if descendant != -1 {
				Descendant_nr = descendant
			}
		case floor := <- chan_order_executed:
			order := Network.NewOrder{floor, 0, 0, 1, 0}
			Network.Udp_send_order_status(order, *Master_IP)
		}
	}
}

func slave_receive_order_from_hw(chan_new_order chan EventManager.Orders) {
	for {
		DriveElevator.Driveelevator_get_new_order(chan_new_order)
	}
}

func slave_mechanical_problem_send_to_master(order Network.NewOrder, Master_IP string) {
	if EventManager.Eventmanager_stop_mechanical_reason() == 1 {
		order.In_progress = 0
		order.Is_executed = 0
		Network.Udp_send_order_status(order, Master_IP)
	}
}
