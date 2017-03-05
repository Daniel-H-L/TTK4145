package Master

import (
	"./DriveElevator"
	"./Network"
)

func Master_queue_receive_order_executed(chan_order_executed chan int, chan_received_msg chan []byte, port_nr int, chan_error chan error) int {
	go Network.Udp_receive_order_executed(chan_order_executed, chan_received_msg, port_nr, chan_error)
	return <-chan_order_executed
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

func master_queue_detect_local_orders_from_slaves(chan_local_order chan LocalOrder, chan_received_msg chan []byte, portNr string, chan_error chan error) LocalOrder {
	go Network.Udp_receive_local_orders(chan_local_order, chan_received_msg, portNr, chan_error)

	err := <-chan_error
	if err != nil {
		//TO SOMETHING
	}

	return <-chan_local_order
}

func master_queue_detect_local_orders() {
	floor, button := DriveElevator.Eventmanager_get_new_order()
	if button != 0 {
		if button == 3 {
			//Inside order
			//Hvordan få tak i floor og button??
			DriveElevator.Eventmanager_add_new_order(floor, button)
		}

	}
}
