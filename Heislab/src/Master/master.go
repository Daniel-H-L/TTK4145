package Master

import (
	"../DriveElevator"
	"../Network"
	"fmt"
	"time"
)

type Slave struct {
	IP         string
	Descendant int
	Last_floor int
	Direction  int
}

type Master struct {
	Slaves map[string]chan bool
	IP     string
}

func Master_init(chan_master chan Master, chan_received_msg chan Network.StandardData, chan_is_alive chan string, portNr string, chan_change_to_slave chan bool, chan_master_ip chan string, chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button, chan_source_ip chan string, chan_set_lights chan [3][4]int, chan_order_executed chan int, chan_state chan int, chan_dir chan int, chan_floor chan int, chan_elev_state chan Network.ElevState, chan_reset chan bool, chan_descendant_nr chan int, chan_slavelist chan map[string]chan bool, chan_network_lights chan [3][4]int, chan_new_network_order chan Network.NewOrder) {
	fmt.Println("Master init...")
	master := Master{}
	master.IP, _ = Network.Udp_get_local_ip()
	fmt.Println("MasterIP:", master.IP)

	chan_master <- master
	chan_master_ip <- master.IP

	backup := Network.Backup{}
	master_queue := &Network.Queue{}
	master_queue.Orders = [3][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	backup.MainQueue = make(map[string]*Network.Queue)
	backup.MainQueue[master.IP] = master_queue

	go Network.Udp_broadcast(master.IP)
	//Network.Udp_send_is_alive("129.241.187.146")
	time.Sleep(50 * time.Millisecond)
	//order := DriveElevator.Button{1,1}

	order2 := Network.NewOrder{1, 1, 0, 0, 0}
	for {

		Network.Udp_send_new_order(order2, "129.241.187.154")
		time.Sleep(500 * time.Microsecond)
	}

	//go Test_drive(chan_new_hw_order, chan_new_master_order, backup, chan_source_ip, chan_set_lights, chan_order_executed, chan_state, chan_dir, chan_floor, chan_master_ip, chan_received_msg, chan_elev_state, portNr, chan_is_alive, chan_reset, chan_descendant_nr, chan_master, chan_slavelist, chan_network_lights, chan_new_network_order)

}

func slave_timer(chan_reset chan bool, chan_timeout chan string, ip string) {
	for {
		select {
		case <-chan_reset:
			continue
		case <-time.After(5 * time.Second):
			chan_timeout <- ip
			return
		}
	}
}

func update_slave_list(chan_is_alive chan string, chan_slavelist chan map[string]chan bool, chan_reset chan bool, chan_descendant chan int) {

	participants := 0
	slavelist := make(map[string]chan bool)
	chan_timeout := make(chan string)
	for {
		select {
		case slave_io := <-chan_is_alive:
			// ny slave legg til i map eller
			// refresh timer
			t := 0
			for s := range slavelist {
				if s == slave_io {
					chan_reset <- true
					t = 1
				}
			}
			if t == 0 {
				participants += 1
				new_slave := Slave{slave_io, participants, -1, -1}
				slavelist[new_slave.IP] = make(chan bool, 1)
				go slave_timer(slavelist[new_slave.IP], chan_timeout, new_slave.IP)
				chan_descendant <- -1
			}
		case ip := <-chan_timeout: //timeout på slave
			// fjern slave
			delete(slavelist, ip)
			participants -= 1
			chan_descendant <- -1

		}
		//send slave list
		chan_slavelist <- slavelist
	}
}

func write_backup_queue(backup Network.Backup, button DriveElevator.Button, source_ip string, chan_set_lights chan [3][4]int) {
	set_lights := [3][4]int{}
	for ip, order := range backup.MainQueue {
		if ip == source_ip {
			fmt.Println("Order:", *order)
			order.Orders[button.Dir][button.Floor] = 1
			//fmt.Println("Order2: ", *order)
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if order.Orders[i][j] == 1 {
					set_lights[i][j] = 1
				} else {
					set_lights[i][j] = 0
				}
			}
		}
	}
	fmt.Println("TURN OFF THE LIGHTS", set_lights)
	chan_set_lights <- set_lights
}

func reset_backup_queue(backup Network.Backup, floor int, chan_set_lights chan [3][4]int) {
	set_lights := [3][4]int{}

	for _, order := range backup.MainQueue {
		for i := 0; i < 3; i++ {
			order.Orders[i][floor] = 0
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if order.Orders[i][j] == 1 {
					set_lights[i][j] = 1
				} else {
					set_lights[i][j] = 0
				}
			}
		}
	}
	chan_set_lights <- set_lights
}

func reset_backup_slaves(backup Network.Backup, slavelist map[string]chan bool) {
	for ip, elev := range backup.MainQueue {
		if _, key := slavelist[ip]; key {
			continue
		} else {
			elev.State = -1
		}
	}
}

func Test_drive(chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button, backup Network.Backup, chan_source_ip chan string, chan_set_lights chan [3][4]int, chan_order_executed chan int, chan_state chan int, chan_dir chan int, chan_floor chan int, chan_master_ip chan string, chan_received_msg chan Network.StandardData, chan_elev_state chan Network.ElevState, portNr string, chan_is_alive chan string, chan_reset chan bool, chan_descendant_nr chan int, chan_master chan Master, chan_slavelist chan map[string]chan bool, chan_network_lights chan [3][4]int, chan_new_network_order chan Network.NewOrder) {
	fmt.Println("Master test drive start... ")
	Master_IP := <-chan_master_ip
	master := <-chan_master
	chan_kill_network := make(chan bool)

	go Network.Udp_receive_standard_data(chan_received_msg, portNr, chan_source_ip, chan_descendant_nr, chan_new_network_order, chan_elev_state, chan_network_lights, chan_kill_network)

	for {
		select {
		case new_hw_order := <-chan_new_hw_order:
			fmt.Println("Master received hw_order")
			chan_new_master_order <- new_hw_order
			elevator := Master_IP
			if new_hw_order.Dir != 2 {
				order := Network.NewOrder{}
				order.Floor = new_hw_order.Floor
				order.Button = new_hw_order.Dir
				elevator = Delegate_order(order, backup)
			}
			write_backup_queue(backup, new_hw_order, elevator, chan_set_lights)
			set_lights := <-chan_set_lights
			Network.Udp_send_set_lights(set_lights, elevator)

		case executed := <-chan_order_executed:
			fmt.Println("Executed... ", executed)
			reset_backup_queue(backup, executed, chan_set_lights)
			//Network.Udp_send_set_lights(chan_set_lights, elevator)

		case state := <-chan_state:
			//fmt.Println("state (master): ", state)
			for ip, elev := range backup.MainQueue {
				if ip == Master_IP {

					elev.State = state
				}
			}
		case dir := <-chan_dir:
			//fmt.Println("dir (master): ", dir)
			for ip, elev := range backup.MainQueue {
				if ip == Master_IP {

					elev.Direction = dir
				}
			}
		case floor := <-chan_floor:
			//fmt.Println("floor (master): ", floor)
			for ip, elev := range backup.MainQueue {
				if ip == Master_IP {

					elev.Floor = floor
				}
			}
		case status := <-chan_elev_state:
			source_ip := <-chan_source_ip
			for elev := range backup.MainQueue {
				if elev == source_ip {
					backup.MainQueue[elev].State = status.State
					backup.MainQueue[elev].Floor = status.Floor
					backup.MainQueue[elev].Direction = status.Direction
				}
			}
		case slavelist := <-chan_slavelist:
			master.Slaves = slavelist
			reset_backup_slaves(backup, slavelist)
		//case nr := <- chan_descendant:
		//send descendant nr to all slaves.

		default:
			//fmt.Println("Running testdriver")

		}
	}

}
