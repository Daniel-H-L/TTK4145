package Master

import (
	"./Network"
)

func Master_queue_receive_order_executed(chan_order_executed chan int, chan_received_msg chan []byte, port_nr int, chan_error chan error) int {
	go Network.Udp_receive_order_executed(chan_order_executed, chan_received_msg, port_nr, chan_error)
	return <-chan_order_executed
}

func Master_queue_decide_elevator() {

}
