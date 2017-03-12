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
	//Last_floor int
	//Direction  int
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

	//go Network.Udp_broadcast(master.IP)
	//Forlsag:
	//broadcastCh <- master.IP

	time.Sleep(50 * time.Millisecond)

	//go Test_drive()

}

func slaveTimer(resetCh chan bool, timeoutCh chan string, ip string) {
	for {
		select {
		case <-resetCh:
			continue
		case <-time.After(5 * time.Second):
			timeoutCh <- ip
			return
		}
	}
}

func updateSlaveList(isAliveCh chan string, slavelistCh chan map[string]chan bool, resetCh chan bool, descendantCh chan int) {

	participants := 0
	slavelist := make(map[string]chan bool)
	timeoutCh := make(chan string)
	for {
		select {
		case slaveIO := <-isAliveCh:
			// ny slave legg til i map eller
			// refresh timer
			t := 0
			for s := range slavelist {
				if s == slaveIO {
					resetCh <- true
					t = 1
				}
			}
			if t == 0 {
				participants += 1
				newSlave := Slave{slaveIO, participants, -1, -1}
				slavelist[newSlave.IP] = make(chan bool, 1)
				go slave_timer(slavelist[newSlave.IP], timeoutCh, newSlave.IP)
				descendantCh <- -1
			}
		case ip := <-timeoutCh: //timeout på slave
			// fjern slave
			delete(slavelist, ip)
			participants -= 1
			descendantCh <- -1

		}
		//send slave list
		slavelistCh <- slavelist
	}
}

func writeBackupQueue(backup map[string]*Network.Backup, button DriveElevator.Button, source_ip string, setLightsCh chan [3][4]int) {
	setLights := [3][4]int{}
	for ip, order := range backup.MainQueue {
		if ip == source_ip {
			fmt.Println("Order:", *order)
			order.Orders[button.Dir][button.Floor] = 1
			//fmt.Println("Order2: ", *order)
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if order.Orders[i][j] == 1 {
					setLights[i][j] = 1
				} else {
					setLights[i][j] = 0
				}
			}
		}
	}
	fmt.Println("TURN OFF THE LIGHTS", set_lights)
	setLightsCh <- setLights
}

func resetBackupQueue(backup map[string]*Network.Backup, floor int, setLightsCh chan [3][4]int) {
	setLights := [3][4]int{}

	for _, order := range backup {
		for i := 0; i < 3; i++ {
			order.Orders[i][floor] = 0
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if order.Orders[i][j] == 1 {
					setLights[i][j] = 1
				} else {
					setLights[i][j] = 0
				}
			}
		}
	}
	setLightsCh <- setLights
}

func resetBackupSlaves(mainBackupCh map[string]*Network.Backup, slavelist map[string]chan bool) {
	mainBackup := <-mainBackupCh
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

	for ip := range mainBackup {
		//Update current elevator state
		mainBackup[ip].Direction = backup[ip].Direction
		mainBackup[ip].Floor = backup[ip].Floor
		mainBackup[ip].State = backup[ip].State

		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if mainBackup[ip].Orders[i][j] == 0 && backup[ip].Orders[i][j] == 1 {
					//New order
					elevator := desideElevator(i, j, mainBackup)
					mainBackup[elevator].Orders[i][j] = 1

				} else if mainBackup[ip].Orders[i][j] == 1 && backup[ip].Orders[i][j] == 0 {
					//Ikke mottatt
					//Send på nytt?
				} else if mainBackup[ip].Orders[i][j] == 2 && backup[ip].Orders[i][j] == 0 {
					mainBackup[ip].Orders[i][j] = 0

				} else if mainBackup[ip].Orders[i][j] == 1 && backup[ip].Orders[i][j] == 1 {
					mainBackup[ip].Orders[i][j] = 2

				} else if mainBackup[ip].State == -1 && backup[ip].Orders[i][j] == 2 {
					//Handle dead elevators
					elevator := desideElevator(i, j, mainBackup)
					mainBackup[elevator].Orders[i][j] = 1
				}
			}
		}
	}
	mainBackupCh <- mainBackup
}

func updateBackupForResurrectedElevator(elevatorIP string, mainBackupCh chan map[string]*Network.Backup, slaveBackup map[string]*Network.Backup) {
	mainBackup := <-mainBackupCh

	slaveBackup = mainBackup

}

func driveElevator(portNr string, masterCh chan Master, newHWOrder chan [3][4]int) {
	fmt.Println("Master test drive start... ")

	newMasterOrderCh := make(chan DriveElevator.Button, 100)

	slavelistCh := make(chan map[string]chan bool)

	for {
		select {
		case backup := <-slaveBackupCh:

		case slavelist := <-slavelistCh:
			master := <-masterCh
			master.Slaves = slavelist
			resetBackupSlaves(backup, slavelist)
		//case nr := <- chan_descendant:
		//send descendant nr to all slaves.

		default:
			//fmt.Println("Running testdriver")

		}
	}

}
