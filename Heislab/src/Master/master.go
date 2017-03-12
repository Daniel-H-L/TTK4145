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

func MasterInit(chan_change_to_slave chan bool, portNr string) {
	fmt.Println("Master init...")
	master := Master{}
	master.IP, _ = Network.Udp_get_local_ip()
	fmt.Println("MasterIP:", master.IP)

	go Network.Udp_broadcast(master.IP)

	time.Sleep(50 * time.Millisecond)

	//go Test_drive()

}

func slaveTimer(chan_reset chan bool, chan_timeout chan string, ip string) {
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

func updateSlaveList(chan_is_alive chan string, chan_slavelist chan map[string]chan bool, chan_reset chan bool, chan_descendant chan int) {

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

func writeBackupQueue(backup map[string]*Network.Backup, button DriveElevator.Button, source_ip string, chan_set_lights chan [3][4]int) {
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

func resetBackupQueue(backup map[string]*Network.Backup, floor int, chan_set_lights chan [3][4]int) {
	set_lights := [3][4]int{}

	for _, order := range backup {
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

func resetBackupSlaves(backup map[string]*Network.Backup, slavelist map[string]chan bool) {
	for ip, elev := range backup {
		if _, key := slavelist[ip]; key {
			continue
		} else {
			elev.State = -1
		}
	}
}

func delegateOrder(backup, mainBackupCh chan map[string]*Backup) {
	mainBackup := <-mainBackupCh

	//Kjør kostfunskjon

	//Update backup
	for ip := range mainBackup {
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if mainBackup[ip].Orders[i][j] == 0 && backup[ip].Orders[i][j] == 1 {
					mainBackup[ip].Orders[i][j] = 1
				} else if mainBackup[ip].Orders[i][j] == 1 && backup[ip].Orders[i][j] == 0 {
					//Ikke mottatt
					//Send på nytt?
				} else if mainBackup[ip].Orders[i][j] == 2 && backup[ip].Orders[i][j] == 0 {
					mainBackup[ip].Orders[i][j] = 0
				} else if mainBackup[ip].Orders[i][j] == 1 && backup[ip].Orders[i][j] == 1 {
					mainBackup[ip].Orders[i][j] = 2
				}
			}
		}
	}
}

func driveElevator(portNr string, masterCh chan Master, newHWOrder chan [3][4]int) {
	fmt.Println("Master test drive start... ")

	chan_kill_network := make(chan bool)

	newHWOrder := make(chan DriveElevator.Button, 100)
	newMasterOrderCh := make(chan DriveElevator.Button, 100)

	slavelistCh := make(chan map[string]chan bool)

	for {
		select {
		case backup := <-slaveBackupCh:

		case slavelist := <-chan_slavelist:
			master := <-chan_master
			master.Slaves = slavelist
			resetBackupSlaves(backup, slavelist)
		//case nr := <- chan_descendant:
		//send descendant nr to all slaves.

		default:
			//fmt.Println("Running testdriver")

		}
	}

}
