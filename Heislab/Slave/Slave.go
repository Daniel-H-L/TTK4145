package Slave

import (
	"./Network"
)

var Master_IP string
var Descendant_nr int

func Slave_init() {
	Master_IP = Slave_listen_for_master()
}

func Slave_send_is_alive() {
	Network.Udp_send_is_alive(Master_IP)
}

func Slave_listen_for_master() string {

}

func Slave_order_executed(order int) {
	Network.Udp_send_order_executed(order)
}

func Slave_send_local_order(order LocalOrder) {
	Network.Udp_send_local_order(order, Master_IP)
}

func Slave_receive_order(chan_new_order chan NewOrder, chan_received_msg chan []byte, port_nr string, chan_error chan error) {
	go Network.Udp_receive_new_order(chan_new_order, chan_received_msg, port_nr, chan_error)
	order := <-chan_new_order
	//Drive_elevator(order)

}

func Slave_to_master() {

}

func Slave_receive_descendant(chan_descendant_nr chan int, chan_received_msg chan []byte, port_nr string, chan_error chan error) {
	go Network.Udp_receive_descendant_nr(chan_descendant_nr, chan_received_msg, port_nr, chan_error)
	Descendant_nr := <-chan_descendant_nr

}
