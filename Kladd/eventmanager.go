package EventManager

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "eventmanager.h"
#include "queue.h"
#include "stdio.h"
*/
import "C"

type Orders struct {
	Floor int
	Dir   int
}

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

func Eventmanager_check_button_signal(chan_new_order chan Orders) {
	var prev_order Orders

	for {
		new_order := Orders{}
		new_order_C := C.check_button_signal()
		new_order.Floor = int(new_order_C.floor)
		new_order.Dir = int(new_order_C.dir)
		if new_order.Dir != -1 && prev_order != new_order{
			prev_order = new_order
			chan_new_order <- new_order
		}
	}
}

func Eventmanager_add_new_order(floor int, button int) {
	C.queue_add_new_order(C.int(floor), C.int(button))
	Statemachine_set_button_lights()
}

func Eventmanager_stop_mechanical_reason() int {
	return int(C.stop_mechanical_reason())
}
