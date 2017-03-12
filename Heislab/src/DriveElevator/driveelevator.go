package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

func Run_elevator(chan_state chan int, chan_dir chan int, chan_floor chan int, chan_order_executed chan int, newHWOrderCh chan Button, chan_new_master_order chan map[string]*Network.Backup, chan_set_lights chan [3][4]int) {
	fmt.Println("Run elev...")
	if EventManager.ElevatorInit() != 1 {
		fmt.Println("Uanble to initialize elevator hardware... \n")
		return
	}

	EventManager.ElevatorSetMotorDir(1)
	for {
		if EventManager.ElevatorGetFloorSensorSignal() != -1 {
			break
		}
		fmt.Println("Floor sensor is -1...")
	}
	EventManager.ElevatorSetMotorDir(0)

	chan_timer := make(chan bool, 1)
	floorSensorCh := make(chan int, 1)
	go InternalQueuePollButtons(newHWOrderCh)
	go InternalQueuePollFloorSensors(floorSensorCh)
	go timer(chan_timer)
	go StatemachineSetLights(seLightsCh)

	timer := time.NewTimer(3 * time.Second)
	timer.Stop()

	for {
		select {
		case new_order := <-chan_new_master_order:
			fmt.Println("NEW ORDER: ", new_order)
			StatemachineButtonPush(new_order, chan_dir, chan_state, timer)
			time.Sleep(50 * time.Millisecond)

		case floor := <-floorSensorCh:
			ElevatorFloor = floor

			//fmt.Println("NEW FLOOR: ", floor)
			EventManager.Elevator_set_floor_indicator(floor)
			StatemachineArrivedAtFloor(floor, chan_order_executed, timer)
			//time.Sleep(50 * time.Millisecond)

		case <-timer.C:
			StatemachineDoorTimeOut(chan_dir, chan_state)

			// case executed := <-chan_order_executed:
			// 	fmt.Println("executed", executed)
		}
	}
}
