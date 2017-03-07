package Master

import (
	"./Network"
	"fmt"
	"time"
)

type Slave struct {
	IP         string
	Alive      bool
	Descendant int
	Last_floor int
	Direction  int
}

type Master struct {
	Slaves       [Slave]time.Time
	IP           string
	Participants int
	Last_floor   int
	Direction    int
}

func (master *Master) master_init() {
	fmt.Println("Master init...")
	master.IP = Network.Udp_get_local_ip()
	master.Participants = 0

	go Network.Udp_broadcast("")
}

func master_detect_slave(chan_rec_msg chan []byte, chan_is_alive chan string, port_nr string, chan_error chan error) {
	go Network.Udp_receive_is_alive(chan_rec_msg, chan_is_alive, port_nr, chan_error)

	master := Master{}
	master.master_init()
	for {
		msg := <-chan_is_alive
		is_updated := false

		if msg != string() {
			for s := range master.Slaves {
				if s.IP == msg {
					master.Slaves[s] = time.Now()
					is_updated := true
				}
			}
			if is_updated == false {
				new_slave := Slave{msg, true, master.Participants + 1}
				master.Slaves[new_slave] = time.Now()
				master.Participants += 1
			}
		}

		const DEADLINE = 1 * time.Second
		descendant := -1
		for slave, last_time := range master.Slaves {
			if time.Since(last_time) > DEADLINE {
				descendant = slave.Descendant
				delete(master.Slaves, slave)
				master.Participants -= 1
			}
		}
		for s := range master.Slaves {
			if s.Descendant > descendant {
				s.Descendant -= 1
				Network.Udp_send_descendant_nr(s.Descendant, s.IP)
			}
		}
	}

}

func master_resend() {

}

func master_order_executed(chan_order_executed chan int, chan_received_msg chan []byte, port_nr int, chan_error chan error) { //Snakker med master_queue-modulen. Feilhåndtering.
	Master_queue_receive_order_executed(chan_order_executed, chan_received_msg, port_nr, chan_error)
}

func (master *Master) Master_send_is_alive() {
	for slave := range master.Slaves {
		Network.Udp_send_is_alive(s.IP)
	}
}

//Må oppdatere direction og floor variablene hos master og slaver.
