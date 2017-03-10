package EventManager

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
#include "stdio.h"
*/
import "C"

//Translate enum from elev.h
type MotorDirection int
type ElevatorButton int

Vi skal bruke simulator

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

func Elevator_init() int {
	return int(C.elev_init())
}

func Elevator_set_motor_dir(direction int) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(direction))
}

func Elevator_set_button_lamp(button int, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func Elevator_set_floor_indicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func Elevator_set_door_open_lamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}

func Elevator_get_button_signal(button int, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func Elevator_get_floor_sensor_signal() int {
	return int(C.elev_get_floor_sensor_signal())
}
