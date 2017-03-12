package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

var ElevatorState int

var MotorDir int

var ElevatorFloor int

const (
	IDLE      = 0
	MOVING    = 1
	DOOR_OPEN = 2
)

const (
	UP   = 1
	DOWN = -1
	STOP = 0
)

func StatemachineArrivedAtFloor(floor int, backupCh chan Network.Backup, timer *time.Timer) {
	switch ElevatorState {
	case MOVING:
		if QueueShouldStop(MotorDir, ElevatorFloor) == 1 {

			EventManager.ElevatorSetMotorDir(0)
			EventManager.ElevatorSetDoorOpenLamp(1)
			timer.Reset(3 * time.Second)
			InternalQueueDeleteOrder(floor, backupCh)

			ElevatorFloor = floor
			ElevatorState = DOOR_OPEN //flytte denne opp?
			backup := <-backupCh
			backup.State = ElevatorState
			backup.Floor = ElevatorFloor

			fmt.Print("order executed HW... ", floor, "\n")

		}
	}
}

func StatemachineDoorTimeOut(chan_dir chan int, chan_state chan int) {
	switch ElevatorState {
	case DOOR_OPEN:
		EventManager.ElevatorSetDoorOpenLamp(0)
		dir := QueueChooseDir()
		//fmt.Println("dir: ", dir)

		if dir != MotorDir {
			chan_dir <- dir
		}

		MotorDir = dir
		EventManager.ElevatorSetMotorDir(dir)

		if dir != 0 {
			ElevatorState = MOVING
			chan_state <- ElevatorState
		} else {
			ElevatorState = IDLE
			chan_state <- ElevatorState
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func StatemachineButtonPush(button Button, backupCh chan Network.Backup, timer *time.Timer) {
	backup := <-backupCh
	switch ElevatorState {
	case IDLE:
		if ElevatorFloor == button.Floor {
			EventManager.ElevatorSetDoorOpenLamp(1)
			timer.Reset(3 * time.Second)
			ElevatorState = DOOR_OPEN
			time.Sleep(50 * time.Millisecond)
		} else {
			backup.Orders[button.Dir][button.Floor] = 1
			dir := QueueChooseDir()
			//fmt.Println("dir: ", dir)

			if dir != MotorDir {
				backup.Direction = dir
			}
			//time.Sleep(50 * time.Millisecond)

			if dir != 0 {
				MotorDir = dir
				EventManager.ElevatorSetMotorDir(dir)
				ElevatorState = MOVING
			}
		}
	case MOVING:
		Internal_queue[button.Dir][button.Floor] = 1
		//fmt.Println("MOVING 2")
		break
	case DOOR_OPEN:
		if ElevatorFloor == button.Floor {
			timer.Reset(3 * time.Second)
			QueueDeleteOrder(ElevatorFloor)
		} else {
			Internal_queue[button.Dir][button.Floor] = 1
		}
	}
	backup.State = ElevatorState
	backupCh <- backup
}

func StatemachineSetLights(setLightsCh chan [3][4]int) {
	for {
		select {
		case setLights := <-setLightsCh:
			fmt.Println("Set lights: ", setLights)
			for i := 0; i < 3; i++ {
				for j := 0; j < 4; j++ {
					if setLights[i][j] == 1 {
						EventManager.ElevatorSetButtonLamp(i, j, 1)
					} else {
						EventManager.ElevatorSetButtonLamp(i, j, 0)
					}
				}
			}
		}
	}
}
