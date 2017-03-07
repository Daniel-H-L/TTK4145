package EventManager

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "eventmanager2.h"
#include "stdio.h"
*/
import "C"

type Orders struct {
	floor     int
	dir       int
	is_inside bool
}

func Eventmanager_new_order_in_empty_queue() int {
	return int(C.new_order_in_empty_queue2())
}

func Eventmanager_arrive_at_floor() {
	C.arrive_at_floor2()
}

func Eventmanager_orders_in_same_floor() {
	C.orders_in_same_floor2()
}

func Eventmanager_door_timeout() int {
	return int(C.door_time_out2())
}

func Eventmanager_check_button_signal() Orders {
	new_order := Orders{}
	new_order_C := C.check_button_signal2()
	new_order.floor = int(new_order_C.floor)
	new_order.dir = int(new_order_C.dir)
	new_order.is_inside = bool(new_order_C.is_inside)

	return new_order
}

func Eventmanager_add_new_order(floor int, button int) {
	C.queue_add_new_order(floor, button)
	Statemachine_set_button_lights()
}
