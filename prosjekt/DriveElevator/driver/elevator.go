package driver

import (
	"fmt"
	"unsafe"
)

import "C"

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
#include "stdio.h"
*/

//Translate enum from elev.h
type MotorDirection int
type ElevatorButton int

const (
	MOTOR_DIR_UP   MotorDirection = 1
	MOTOR_DIR_DOWN MotorDirection = -1
	MOTOR_DIR_STOP MotorDirection = 0
)

const (
	ELEV_BUTTON_UP      ElevatorButton = 0
	ELEV_BUTTON_DOWN    ElevatorButton = 1
	ELEV_BUTTON_COMMAND ElevatorButton = 2
)

func elevator_init() int {
	return int(C.elev_init())
}

func elevator_set_motor_dir(direction MotorDirection) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(direction))
}

func elevator_set_button_lamp(button ElevatorButton, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func elevator_set_floor_indicator(int floor) {
	C.elev_set_floor_indicator(C.int(floor))
}

func elevator_set_door_open_lamp(int value) {
	C.elev_set_door_open_lamp(C.int(value))
}

func elevator_get_button_signal(button ElevatorButton, floor int) {
	return int(C.elev_get_button_signal(C.elev_get_button_signal(button), C.int(floor)))
}

func elevator_get_floor_sensor_signal() {
	return int(C.elev_get_floor_sensor_signal())
}