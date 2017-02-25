package Master

import (
	"./Network"
	"fmt"
)

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

type StandardData struct {
	IP     string
	Msg_ID string
	//is_master         bool
	Is_alive       bool //overflødig?
	Order_executed int
	Elevator_nr    int //overflødig?
	Descendant_nr  int
	New_order      NewOrder
	Local_order    LocalOrder
	//last_floor Floor
	//dir
	//backup ???
}

// type Participants bool

// const (
// 	MASTER Participants = True
// 	SLAVE1 Participants = False
// 	SLAVE2 Participants = False
// )

type Participants string

// type Is_alive bool

const (
	MASTER Participants = ""
	SLAVE1 Participants = ""
	SLAVE2 Participants = ""
)

// const (
// 	MASTER Is_alive = False
// 	SLAVE1 Is_alive = False
// 	SLAVE2 Is_alive = False
// )

func master_init() {
	fmt.Println("Master init...")

	chan_is_alive := make(chan string, 1)
	chan_received_msg := make(chan []byte, 1)
	//chan_is_master := make(chan bool, 1)
	chan_order_executed := make(chan int, 1)
	chan_descendant_nr := make(chan int, 1)
	chan_new_order := make(chan NewOrder, 1)
	chan_local_order := make(chan LocalOrder, 1)
	chan_error := make(chan error, 1)

	//broadcast IP
	Network.Udp_broadcast("", 1)
	//decide descendantnr
}

func master_change_to_slave() {
	//
}

func master_decide_descendant_nr() {

}

func master_detect_slave_alive(chan_rec_msg chan []byte, chan_is_alive chan string, port_nr string, chan_error chan error) {
	go Network.Udp_receive_is_alive(chan_rec_msg, chan_is_alive, port_nr, chan_error)

	err := <-chan_err
	if err != nil {
		//DO SOMETHING
	}
	slave_ip := <-chan_is_alive
	if SLAVE1 == "" {
		SLAVE1 = slave_ip
	} else if SLAVE2 == "" {
		SLAVE2 = slave_ip
	}

	for {
		//Sleep
		//Sjekk kont for levende slaver. dersom noen e døde --> endre Is_alive variabelen.
	}

}
