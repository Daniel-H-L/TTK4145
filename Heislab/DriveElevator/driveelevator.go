package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

func Driveelevator_get_new_order(chan_new_order chan Orders) {
	order := EventManager.Eventmanager_check_button_signal()
	if order.dir != -1 {
		//New order is detected
		chan_new_order <- order
	}
}

func Driveelevator_order_executed(chan_order_executed chan int) {
	floor := Statemachine_send_deleted_order()
	if floor != -1 {
		chan_order_executed <- floor
	}
}

func Driveelevator_state_update(chan_state chan int, chan_dir chan int, chan_floor chan int) {

}

func Driveelevator_check_inputs() {

}

func Driveelevator_send_outputs() {

}

func Run_elevator(chan_state chan int, chan_dir chan int, chan_floor chan int, chan_order_executed chan int, chan_new_order chan Orders) { //need new name
	if EventManager.Elevator_init() != 1 {
		fmt.Println("Uanble to initialize elevator hardware... \n")
		return
	}
	for EventManager.Elevator_get_floor_sensor_signal() == -1 {
		EventManager.Statemachine_set_state_and_dir(EventManager.MOVING, EventManager.MOTOR_DIR_UP)
		chan_state <- int(EventManager.MOVING)
		chan_dir <- int(EventManager.MOTOR_DIR_UP)
	}

	EventManager.Statemachine_set_state_and_dir(EventManager.IDLE, EventManager.MOTOR_DIR_STOP)
	chan_current_state <- int(EventManager.MOVING)
	chan_current_dir <- int(EventManager.MOTOR_DIR_UP)

	for {
		go Driveelevator_get_new_order(chan_new_order)

		chan_floor <- int(EventManager.Statemachine_set_current_floor())

		EventManager.Statemachine_set_button_lights()
		EventManager.Eventmanager_new_order_in_empty_queue()
		EventManager.Eventmanager_arrive_at_floor()
		EventManager.Eventmanager_door_timeout()
		EventManager.Eventmanager_orders_in_same_floor()
		time.Sleep(400 * time.Millisecond)
	}
}
