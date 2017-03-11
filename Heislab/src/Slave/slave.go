package Slave

import (
	"../DriveElevator"
	//"../DriveElevator/EventManager"
	"../Network"
	"fmt"
	"time"
)

var Descendant_nr int

func Slave_init(portNr string, chan_error chan error, chan_change_to_master chan bool, Master_IP *string, state int, chan_new_network_order chan Network.NewOrder, chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button, chan_order_executed chan int, chan_set_lights chan [3][4]int, chan_network_lights chan [3][4]int, port_nr string, chan_source_ip chan string, chan_descendant_nr chan int, chan_new_order chan Network.NewOrder, chan_elev_state chan Network.ElevState, chan_kill chan bool, chan_kill2 chan bool) {
	*Master_IP = ""

	chan_received_msg := make(chan Network.StandardData, 1)

	go Listen_for_master(chan_received_msg, portNr, chan_error, chan_change_to_master, chan_kill2)
	//go Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr)
	fmt.Println("In slave init...")
	//local_ip, _ := Network.Udp_get_local_ip()

	go Slave_test_drive(chan_new_hw_order, chan_new_master_order, chan_order_executed, chan_set_lights, chan_network_lights, Master_IP, chan_received_msg, port_nr, chan_source_ip, chan_descendant_nr, chan_new_order, chan_elev_state, chan_new_network_order, chan_kill)
}

func Listen_for_master(chan_received_msg chan Network.StandardData, portNr string, chan_error chan error, chan_change_to_master chan bool, chan_kill chan bool) {
	kill_sub_routine := make(chan bool, 1)
	chan_is_alive := make(chan string, 1)

	Descendant_nr = 1
	go Network.Udp_receive_is_alive(chan_received_msg, chan_is_alive, portNr, kill_sub_routine)

	for {
		fmt.Println("Listen for master... I hope someone is broadcasting... So exited!")
		select {
		case ip := <-chan_is_alive:
			fmt.Println("Master detetcted", ip)
			//chan_master_ip <- ip
		case <-time.After(4 * time.Second):
			fmt.Println("Fikk ikke noe på chan_is_alive så derfor er vi her")
			if Descendant_nr == 1 {
				chan_change_to_master <- true

			} else if Descendant_nr != 1 {
				time.Sleep(5 * time.Second)
				Descendant_nr -= 1
			}
		case <-chan_kill:
			fmt.Println("Listen for master received kill")
			kill_sub_routine <- true
			return
		}
		fmt.Println("still in for loop...")
		//time.Sleep(20 * time.Millisecond)
	}
}

func Slave_test_drive(chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button, chan_order_executed chan int, chan_set_lights chan [3][4]int, chan_network_lights chan [3][4]int, Master_IP *string, chan_received_msg chan Network.StandardData, port_nr string, chan_source_ip chan string, chan_descendant_nr chan int, chan_new_order chan Network.NewOrder, chan_elev_state chan Network.ElevState, chan_new_network_order chan Network.NewOrder, chan_kill chan bool) {

	chan_kill_network := make(chan bool)
	go Network.Udp_receive_standard_data(chan_received_msg, port_nr, chan_source_ip, chan_descendant_nr, chan_new_order, chan_elev_state, chan_network_lights, chan_kill_network)
	for {
		select {
		case new_hw_order := <-chan_new_hw_order:
			new_order := Network.NewOrder{new_hw_order.Floor, new_hw_order.Dir, 1, 0, 0}
			Network.Udp_send_new_order(new_order, *Master_IP)

		case new_master_order := <-chan_new_order:
			fmt.Println("received order", new_master_order)
			button := DriveElevator.Button{new_master_order.Floor, new_master_order.Button}
			chan_new_master_order <- button

		case executed := <-chan_order_executed:
			order := Network.NewOrder{executed, 0, 0, 1, 0}
			Network.Udp_send_order_status(order, *Master_IP)

		case set_lights := <-chan_network_lights:
			chan_set_lights <- set_lights
		case kill := <-chan_kill:
			fmt.Println("Killed test drive")
			chan_kill_network <- kill
			return
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
