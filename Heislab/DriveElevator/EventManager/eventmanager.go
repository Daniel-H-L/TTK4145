package EventManager

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "eventmanager.h"
#include "stdio.h"
*/
import "C"

func Eventmanager_new_order_in_empty_queue() int {
	return int(C.new_order_in_empty_queue())
}

func Eventmanager_arrive_at_floor() {
	C.arrive_at_floor()
}

func Eventmanager_orders_in_same_floor() {
	C.orders_in_same_floor()
}

func Eventmanager_door_timeout() int {
	return int(C.door_time_out())
}

func Eventmanager_check_button_signal() int {
	return int(C.check_button_signal())
}
