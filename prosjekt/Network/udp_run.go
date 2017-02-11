package Network

import (
	"fmt"
	"net"
)

type StandardData struct {
	IP     string
	msg_ID string
	//is_master         bool
	is_alive       bool //overflødig?
	order_executed int
	elevator_nr    int //overflødig?
	descendant_nr  int
	new_order      NewOrder
	local_order    LocalOrder
	//last_floor Floor
	//dir
	//backup ???
}

type NewOrder struct {
	floor     int
	direction int //endre til C-definert variabeltype
	priority  int
	order_nr  int
}

type LocalOrder struct {
	is_inside_order bool
	floor           int //endre til C-typen
	direction       int //endre til C-typen
}

func Run_network() {
	chan_udp_bcast := make(chan string, 1)
	chan_is_alive := make(chan string)
	chan_received_msg := make(chan []byte)
	chan_is_master := make(chan bool, 1)
	chan_order_executed := make(chan int)
	chan_descendant_nr := make(chan int)
	chan_new_order := make(chan NewOrder)
	chan_local_order := make(chan LocalOrder)
}
