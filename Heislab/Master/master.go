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
	for {
		Network.Udp_broadcast(master.IP)
		time.Sleep(50 * time.Millisecond)
	}
}

func Master_detect_slave(chan_rec_msg chan []byte, chan_is_alive chan string, port_nr string, chan_error chan error) {
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

func Master_drive_elevator(backup *Backup, chan_new_order chan NewOrder, chan_elev_order chan Driveelevator.Orders, chan_source_ip chan string) {
	go Network.Udp_receive_new_orders(chan_new_order, chan_received_msg, portNr, chan_error, chan_source_ip)
	go DriveElevator.Driveelevator_get_new_order(chan_elev_order)
	go Network.Udp_receive_order_status(chan_order_executed, chan_received_msg, port_nr, chan_error, chan_source_ip)

	for {
		select {
		case new_slave_order := <-chan_new_order:
			source := <-chan_source_ip

			if new_slave_order.direction == 2 { //inside order
				for order := range *backup.MainQueue {
					if *backup.MainQueue[order] == source {
						order.Orders[new_slave_order.direction][new_order.floor] = 1
					}
				}
			} else {
				elevator := master.Master_queue_delegate_order(new_slave_order)
				if elevator != Network.Udp_get_local_ip() {
					Network.Udp_send_new_order(new_slave_order, elevator)
				} else {
					DriveElevator.Eventmanager_add_new_order(new_slave_order.floor, new_slave_order.direction)
				}
			}
		case new_hw_order := <-chan_elev_order:
			elevator := Network.Udp_get_local_ip()

			if new_hw_order.direction == 2 {
				DriveElevator.Eventmanager_add_new_order(new_hw_order.floor, new_hw_order.dir)
			} else {
				order := Network.NewOrder{new_hw_order.floor, new_hw_order.dir, 0, 1, 0, 0}
				elevator = *master.Master_queue_delegate_order(order)

				if elevator != Network.Udp_get_local_ip() {
					Network.Udp_send_new_order(order, elevator)
				} else {
					DriveElevator.Eventmanager_add_new_order(new_hw_order.floor, new_hw_order.dir)
				}
			}
			//add order to main queue
			for order := range *backup.MainQueue {
				if *backup.MainQueue[order] == elevator {
					order.Orders[new_hw_order.dir][new_hw_order.floor] = 1
				}
			}
		case executed := chan_order_executed:
			source_ip := <-chan_source_ip
			if executed.is_executed == true {
				floor := executed.floor
				//slette fra hovedkø med floor og alle buttons.
				for order := range *backup.MainQueue {
					if backup.MainQueue[order] == source_ip {
						for i = 0; i < 3; i++ {
							order.Orders[i][floor] = 0
						}
					}
				}
			}
		}
	}

}

//Må oppdatere direction og floor variablene hos master og slaver.
