package EventManager

import (
	"./driver"
	"fmt"
)
import "C"

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "statemachine.h"
#include "stdio.h"
*/

type ElevatorState int

const (
	IDLE   ElevatorState = 0
	MOVING ElevatorState = 1
	STOP   ElevatorState = 2
)

func Statemachine_set_current_floor() int {
	return C.int(C.set_current_floor())
}

func Statemachine_set_button_lights() {
	C.set_button_signal()
}

func Statemachine_set_state_and_dir(state ElevatorState, dir driver.MotorDirection) {
	C.set_state_and_dir
}

func Statemachine_arrived_at_floor(floor int) int {
	return C.int(C.arrived_at_floor())
}
