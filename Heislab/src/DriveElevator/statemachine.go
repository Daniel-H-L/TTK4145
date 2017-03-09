package DriveElevator

import (
	"./EventManager"
	"fmt"
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

func Statemachine_arrived_at_floor(floor int) {
	switch ElevatorState {
	case MOVING:
		if Internal_queue_should_stop(MotorDir, ElevatorFloor) == 1 {
			EventManager.Elevator_set_motor_dir(0)
			EventManager.Elevator_set_door_open_lamp(1)
			timer_start()
			Internal_queue_delete_order(floor)
			ElevatorFloor = floor
			//si i fra til master at ordre i denne etasjen er utført
			ElevatorState = DOOR_OPEN
		}
	}
}

func Statemachine_door_time_out() {
	switch ElevatorState {
	case DOOR_OPEN:
		EventManager.Elevator_set_door_open_lamp(0)
		dir := Internal_queue_choose_dir()
		fmt.Println("dir: ", dir)
		MotorDir = dir
		EventManager.Elevator_set_motor_dir(dir)
		if dir != 0 {
			ElevatorState = MOVING
		} else {
			ElevatorState = IDLE
		}
	}
}

func Statemachine_button_push(button Button) {
	switch ElevatorState {
	case IDLE:
		if ElevatorFloor == button.floor {
			EventManager.Elevator_set_door_open_lamp(1)
			timer_start()
			ElevatorState = DOOR_OPEN
		} else {
			Internal_queue[button.dir][button.floor] = 1
			dir := Internal_queue_choose_dir()
			fmt.Println("dir: ", dir)
			if dir != 0 {
				MotorDir = dir
				EventManager.Elevator_set_motor_dir(dir)
				ElevatorState = MOVING
			}
		}
	case MOVING:
		Internal_queue[button.dir][button.floor] = 1
		break
	case DOOR_OPEN:
		if ElevatorFloor == button.floor {

		} else {
			Internal_queue[button.dir][button.floor] = 1
		}
	}
}
