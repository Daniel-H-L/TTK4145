package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

type Button struct {
	Dir   int
	Floor int
}

func QueueAddNewOrder(backup Network.Backup) {
	Queue[dir][floor] = 1
}

func QueueDeleteOrder(currentFloor int, backup Network.Backup) {
	for i := 0; i < 3; i++ {
		backup.Orders[i][currentFloor] = 0
	}
}

func QueueDeleteQueue() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			Queue[i][j] = 0
		}
	}
}

func QueueCheckOrdersAbove(currentFloor int) int {
	for i := currentFloor + 1; i < 4; i++ {
		if Queue[0][i] == 1 || Queue[1][i] == 1 || Queue[2][i] == 1 {
			return 1
		}
	}
	return 0
}

func QueueCheckOrdersBlow(currentFloor int) int {
	for i := currentFloor - 1; i > -1; i-- {
		if Queue[0][i] == 1 || Queue[1][i] == 1 || Queue[2][i] == 1 {
			return 1
		}
	}
	return 0
}

func Internal_queue_should_stop(dir int, currentFloor int) int {
	if Queue[2][currentFloor] == 1 {
		return 1
	}
	if dir == UP {
		if Queue[0][currentFloor] == 1 || QueueCheckOrdersAbove(currentFloor) == 0 {
			return 1
		}
	}
	if dir == DOWN {
		if Queue[1][currentFloor] == 1 || QueueCheckOrdersBelow(currentFloor) == 0 {
			return 1
		}
	}
	if dir == STOP {
		if Queue[0][currentFloor] == 1 || Queue[1][currentFloor] == 1 || Queue[2][currentFloor] == 1 {
			return 1
		}
	}
	return 0
}

func QueueChooseDir() int {
	switch MotorDir {
	case UP:
		if QueueCheckOrdersAbove(ElevatorFloor) == 1 {
			return UP
		} else if QueueCheckOrdersBelow(ElevatorFloor) == 1 {
			return DOWN
		} else {
			return STOP
		}
	case DOWN:
		if QueueCheckOrdersBelow(ElevatorFloor) == 1 {
			return DOWN
		} else if QueueCheckOrdersAbove(ElevatorFloor) == 1 {
			return UP
		} else {
			return STOP
		}

	case STOP:
		if QueueCheckOrdersAbove(ElevatorFloor) == 1 {
			return UP
		} else if QueueCheckOrdersBelow(ElevatorFloor) == 1 {
			return DOWN
		} else {
			return STOP
		}
	default:
		return STOP
	}

}

func QueuePollButtons(newHWOrderCh chan Button) {

	var buttonStatus [3][4]bool
	for {
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if !(i == 1 && j == 0) && !(i == 0 && j == 3) {
					if EventManager.ElevatorGetButtonSignal(i, j) == 1 && !buttonStatus[i][j] {
						button := Button{i, j}
						buttonStatus[i][j] = true
						newHWOrderCh <- button
					} else if EventManager.ElevatorGetButtonSignal(i, j) == 0 {
						buttonStatus[i][j] = false
					}
				}
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func QueuePollFloorSensors(floorSensorCh chan int) {
	var prevFloor int
	for {
		floor := EventManager.ElevatorGetFloorSensorSignal()
		if floor != -1 && prevFloor != floor {
			floorSensorCh <- floor
			fmt.Println("floor sensor...", floor)
			prevFloor = floor
		}
		time.Sleep(50 * time.Millisecond)
	}
}
