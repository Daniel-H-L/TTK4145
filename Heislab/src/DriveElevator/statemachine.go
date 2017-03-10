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

func Statemachine_arrived_at_floor(floor int, chan_order_executed chan int, timer *time.Timer) {
	switch ElevatorState {
	case MOVING:
		fmt.Println("MOVING 1")
		if Internal_queue_should_stop(MotorDir, ElevatorFloor) == 1 {
			fmt.Println("STOP!")
			EventManager.Elevator_set_motor_dir(0)
			EventManager.Elevator_set_door_open_lamp(1)
			timer.Reset(3 * time.Second)
			Internal_queue_delete_order(floor)
			fmt.Println("Test1")
			ElevatorFloor = floor
			ElevatorState = DOOR_OPEN //flytte denne opp?
			chan_order_executed <- floor

		}
	}
}

func Statemachine_door_time_out(chan_dir chan int, chan_state chan int) {
	switch ElevatorState {
	case DOOR_OPEN:
		EventManager.Elevator_set_door_open_lamp(0)
		dir := Internal_queue_choose_dir()
		fmt.Println("dir: ", dir)

		if dir != MotorDir {
			chan_dir <- dir
		}

		MotorDir = dir
		EventManager.Elevator_set_motor_dir(dir)

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

func Statemachine_button_push(button Button, chan_dir chan int, chan_state chan int, timer *time.Timer) {
	switch ElevatorState {
	case IDLE:
		if ElevatorFloor == button.Floor {
			EventManager.Elevator_set_door_open_lamp(1)
			timer.Reset(3 * time.Second)
			ElevatorState = DOOR_OPEN
			chan_state <- ElevatorState
			time.Sleep(50 * time.Millisecond)
		} else {
			Internal_queue[button.Dir][button.Floor] = 1
			dir := Internal_queue_choose_dir()
			fmt.Println("dir: ", dir)

			if dir != MotorDir {
				chan_dir <- dir
			}
			//time.Sleep(50 * time.Millisecond)

			if dir != 0 {
				MotorDir = dir
				EventManager.Elevator_set_motor_dir(dir)
				ElevatorState = MOVING
				chan_state <- ElevatorState
			}
		}
	case MOVING:
		Internal_queue[button.Dir][button.Floor] = 1
		fmt.Println("MOVING 2")
		break
	case DOOR_OPEN:
		if ElevatorFloor == button.Floor {
			timer.Reset(3 * time.Second)
			Internal_queue_delete_order(ElevatorFloor)
		} else {
			Internal_queue[button.Dir][button.Floor] = 1
		}
	}
}

func Statemachine_set_lights(chan_set_lights chan [3][4]int) {
	for {
		select {
		case set_lights := <-chan_set_lights:
			fmt.Println("Set lights: ", set_lights)
			for i := 0; i < 3; i++ {
				for j := 0; j < 4; j++ {
					if set_lights[i][j] == 1 {
						EventManager.Elevator_set_button_lamp(i, j, 1)
					} else {
						EventManager.Elevator_set_button_lamp(i, j, 0)
					}
				}
			}
		}
	}
}
