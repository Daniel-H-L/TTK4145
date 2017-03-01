package Master

import (
	"./Network"
	"fmt"
	"time"
)

type NewOrder struct {
	floor     int
	direction int //endre til C-definert variabeltype
	//priority  int
	order_nr int
}

type LocalOrder struct {
	is_inside_order bool
	floor           int
	direction       int //endre til C-typen
}

type StandardData struct {
	IP             string
	Msg_ID         string
	Is_alive       bool
	Order_executed int
	Descendant_nr  int
	New_order      NewOrder
	Local_order    LocalOrder
	Last_floor     int
	//dir
	//backup ???
}

type Is_alive bool

type Elevator struct {
	Is_master    bool
	IP           string
	Alive        Is_alive
	Participants int
	Descendant   int
}

var Master Elevator
var Slave1 Elevator
var Slave2 Elevator

func master_init() {
	fmt.Println("Master init...")
	Master.Is_master = true
	Master.IP, _ = Network.Udp_get_local_ip()
	Master.Participants = 0

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
	//detect slaves
}

func master_decide_descendant_nr() { //Kan denne heller legges til i detect_slave_alive???
	if Master.Participants == 1 {
		if Slave1.Alive == true {
			Slave1.Descendant = 1
		} else {
			Slave2.Descendant = 1
		}
	} else if Master.Participants == 2 {
		Slave1.Descendant = 1
		Slave2.Descendant = 2
	} else {
		Slave1.Descendant = 0
		Slave2.Descendant = 0
	}

}

func master_detect_slave_alive(chan_rec_msg chan []byte, chan_is_alive chan string, port_nr string, chan_error chan error) {
	go Network.Udp_receive_is_alive(chan_rec_msg, chan_is_alive, port_nr, chan_error)

	err := <-chan_err
	if err != nil {
		//DO SOMETHING
	}
	slave_ip := <-chan_is_alive
	switch Master.Participants {
	case 0:
		Slave1.Alive = true
		Slave1.IP = slave_ip
		Master.Participants += 1
	case 1:
		if Slave1.Alive == true {
			Slave2.Alive = true
			Slave2.IP = slave_ip
			Master.Participants += 1
		} else {
			Slave1.Alive = true
			Slave1.IP = slave_ip
			Master.Participants += 1
		}
	case 2:
		Slave1.Alive = true
		Slave2.Alive = true
	}

	for {
		//Sleep
		time.Sleep(5 * time.Second)
		//Sjekk kont for levende slaver. dersom noen e døde --> endre Alive variabelen.

		//Master.Participants -= 1
	}

}

func master_resend() {

}

func master_order_executed() { //Snakker med master_queue-modulen. Feilhondtering.

}

func Master_send_is_alive() {
	Network.Udp_send_is_alive(Slave1.IP)
	Network.Udp_send_is_alive(Slave2.IP)
}
