package Slave

import (
	"../DriveElevator"
	//"../DriveElevator/EventManager"
	"../Network"
	"fmt"
	"time"
)

var DescendantNr int

func SlaveInit(portNr string, chan_change_to_master chan bool) {
	Master_IP := ""

	go Listen_for_master(portNr, chan_change_to_master)

	fmt.Println("In slave init...")

	go Slave_test_drive(masterIPCh, portNr, setNetworkLightsCh, setLightsCh, newMasterOrderCh)
}

func ListenForMaster(portNr string, changeToMasterCh bool, isAliveCh chan string) {

	Descendant_nr = 1

	for {
		select {
		case ip := <-isAliveCh:
			fmt.Println("Master detetcted", ip)
			//chan_master_ip <- ip
		case <-time.After(4 * time.Second):
			if DescendantNr == 1 {
				changeToMasterCh <- true

			} else if Descendant_nr != 1 {
				time.Sleep(5 * time.Second)
				DescendantNr -= 1
			}
		}
	}
}

func updateBackup(floor int, dir int, slaveBackupCh chan map[string]*Network.Backup) {
	localIP, _ := Network.Udp_get_local_ip()

	slaveOrders := <-slaveBackupCh

	slaveOrders[localIP].Orders[dir][floor] = 1

	slaveBackupCh <- slaveOrders
}

func udpateHWOrder(backup map[string]*Network.Backup, slaveBackupCh chan map[string]*Network.Backup, newMasterOrderCh chan [3][4]int) {
	localIP, _ := Network.Udp_get_local_ip()

	var order [3][4]int
	slaveBackup := <-slaveBackupCh

	for o := range backup {
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {

				if backup[o].Orders[i][j] == 1 {
					if o == localIP {
						order[i][j] = 1
					}
					//ACK for order = 2
					slaveBackup[o].Orders[i][j] = 2

				} else if backup[o].Orders[i][j] == 0 {
					if o == localIP {
						order[i][j] = 0
					}
					slaveBackup[o].Orders[i][j] = 0
				}
			}
		}
	}
	newMasterOrderCh <- order
	slaveBackupCh <- slaveBackup
}

func driveElevator(masterIPCh chan string, portNr string, setNetworkLightsCh chan [3][4]int, setLightsCh chan [3][4]int, newMasterOrderCh chan [3][4]int) {
	fmt.Println("Slave test drive...")

	for {
		select {
		case newHWOrder := <-newHWOrderCh:
			updateBackup(newHWOrder.Floor, newHWOrder.Dir, slaveBackupCh)

		case backup := <-masterBackupCh:
			updateHWOrder(backup, slaveBackupCh, newMasterOrderCh)

		case lights := <-setNetworkLightsCh:
			setLightsCh <- lights

		case nr := <-DescendantNrCh:
			DescendantNr += nr
		}
	}
}
