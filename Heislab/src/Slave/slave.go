package Slave

import (
	"../DriveElevator"
	//"../DriveElevator/EventManager"
	"../Network"
	"fmt"
	"time"
)

var Descendant_nr int

func Slave_init(chan_received_msg chan []byte, chan_master_bcast chan string, portNr string, chan_error chan error, chan_change_to_master chan bool, Master_IP *string, state int, chan_new_network_order chan Network.NewOrder, chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button, chan_order_executed chan int, chan_set_lights chan [3][4]int, chan_network_lights chan [3][4]int, port_nr string) {
	*Master_IP = ""
	go Listen_for_master(chan_received_msg, chan_master_bcast, portNr, chan_error, chan_change_to_master)
	fmt.Println("In slave init...")
	//local_ip, _ := Network.Udp_get_local_ip()

	go Slave_test_drive(chan_new_network_order, chan_new_hw_order, chan_new_master_order, chan_order_executed, chan_set_lights, chan_network_lights, Master_IP, chan_received_msg, port_nr)
}

func Listen_for_master(chan_received_msg chan []byte, chan_is_alive chan string, portNr string, chan_error chan error, chan_change_to_master chan bool) {
	Descendant_nr = 1
	go Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr)
	for {
		select {
		case ip := <-chan_is_alive:
			fmt.Println("Master detetcted", ip)
			//chan_master_ip <- ip
		case <-time.After(3 * time.Second):
			if Descendant_nr == 1 {
				chan_change_to_master <- true
				return

			} else if Descendant_nr != 1 {
				time.Sleep(5 * time.Second)
				Descendant_nr -= 1
			}
		}
	}
}

func Slave_test_drive(chan_new_network_order chan Network.NewOrder, chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button, chan_order_executed chan int, chan_set_lights chan [3][4]int, chan_network_lights chan [3][4]int, Master_IP *string, chan_received_msg chan []byte, port_nr string) {
	go Network.Udp_receive_set_lights(chan_network_lights, chan_received_msg, port_nr)
	for {
		select {
		case new_hw_order := <-chan_new_hw_order:
			new_order := Network.NewOrder{new_hw_order.Floor, new_hw_order.Dir, 1, 0, 0}
			Network.Udp_send_new_order(new_order, *Master_IP)

		case new_master_order := <-chan_new_network_order:
			button := DriveElevator.Button{new_master_order.Floor, new_master_order.Button}
			chan_new_master_order <- button

		case executed := <-chan_order_executed:
			order := Network.NewOrder{executed, 0, 0, 1, 0}
			Network.Udp_send_order_status(order, *Master_IP)

		case set_lights := <-chan_network_lights:
			chan_set_lights <- set_lights
		}
	}
}

/*
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
		case descendant := <-chan_descendant_nr:
			if descendant != -1 {
				Descendant_nr = descendant
			}
		case floor := <-chan_order_executed:
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
}*/
