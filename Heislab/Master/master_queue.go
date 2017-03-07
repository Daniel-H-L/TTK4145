package Master

import (
	"./DriveElevator"
	"./Network"
)

func Master_queue_receive_order_status(chan_order_executed chan NewOrder, chan_received_msg chan []byte, port_nr int, chan_error chan error) {
	go Network.Udp_receive_order_status(chan_order_executed, chan_received_msg, port_nr, chan_error)
	err <- chan_error
	if err == nil {
		order := chan_order_executed
		if order.is_executed == true {
			floor := order.floor
			//slette fra hovedkø med floor og alle buttons.
		}
	}
}

func (master *Master) Master_queue_delegate_order(order NewOrder) string {
	for s := range master.Slaves {
		switch order.direction {
		case 1: //UP
			if s.Direction == order.direction || s.Direction == 0 {
				if s.Last_floor < order.floor {
					return s.IP
				}
			} else if master.Direction == order.direction || master.Direction == 0 {
				if master.Last_floor < order.floor {
					return master.IP
				}
			} else {

			}
		case -1: //DOWN
		}
	}
}

func master_queue_receive_new_order_slaves(master *Master, chan_new_order chan NewOrder, chan_received_msg chan []byte, portNr string, chan_error chan error) NewOrder {
	go Network.Udp_receive_new_orders(chan_new_order, chan_received_msg, portNr, chan_error)

	err := <-chan_error
	if err == nil {
		order := <-chan_new_order
		if order.is_inside == true {
			//add to main queue
		} else {
			elevator := master.Master_queue_delegate_order(order)
			if elevator != Network.Udp_get_local_ip() {
				Network.Udp_send_new_order(order, elevator)
			} else {
				DriveElevator.Eventmanager_add_new_order(new_order.floor, new_order.dir)
			}
		}
	}
}

func master_queue_receive_new_order_hw(master *Master, chan_new_order chan Driveelevator.Orders) {
	DriveElevator.Driveelevator_get_new_order(chan_new_order)
	new_order := <-chan_new_order

	if new_order.is_inside == true {
		DriveElevator.Eventmanager_add_new_order(new_order.floor, new_order.dir)
	} else {
		order := Network.NewOrder{new_order.floor, new_order.dir, 0, 1, 0, 0}
		elevator := master.Master_queue_delegate_order(order)

		if elevator != Network.Udp_get_local_ip() {
			Network.Udp_send_new_order(order, elevator)
		} else {
			DriveElevator.Eventmanager_add_new_order(new_order.floor, new_order.dir)
		}
	}
	//add to main queue
}
