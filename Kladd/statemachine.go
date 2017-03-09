package EventManager

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "statemachine.h"
#include "stdio.h"

import "C"

var ElevatorState int

const (
	IDLE   = 0
	MOVING = 1
	STOP   = 2
)

func Statemachine_set_current_floor() int {
	return int(C.set_current_floor())
}

func Statemachine_set_button_lights() {
	C.set_button_lights()
}

func Statemachine_set_state_and_dir(state ElevatorState, dir MotorDirection) {
	C.set_state_and_dir(C.state_t(state), C.dir_t(dir))
}

func Statemachine_arrived_at_floor(floor int) int {
	return int(C.arrived_floor(C.int(floor)))
}
*/