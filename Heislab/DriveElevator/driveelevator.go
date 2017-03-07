package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

func Driveelevator_get_new_order(chan_new_order chan Orders) {
	order := EventManager.Eventmanager_check_button_signal()
	if order.dir != 2 {
		//New order is detected
		chan_new_order <- order
	}
}

func run_elevator() { //need new name
	chan_current_floor := make(chan int, 1)
	chan_order_executed := make(chan int, 1)
	chan_current_dir := make(chan int, 1)
	chan_current_state := make(chan int, 1)

	if EventManager.Elevator_init() != 1 {
		fmt.Println("Uanble to initialize elevator hardware... \n")
		return
	}
	for EventManager.Elevator_get_floor_sensor_signal() == -1 {
		EventManager.Statemachine_set_state_and_dir(EventManager.MOVING, EventManager.MOTOR_DIR_UP)

		chan_current_state <- int(EventManager.MOVING)
		chan_current_dir <- int(EventManager.MOTOR_DIR_UP)
	}

	EventManager.Statemachine_set_state_and_dir(EventManager.IDLE, EventManager.MOTOR_DIR_STOP)
	chan_current_state <- int(EventManager.MOVING)
	chan_current_dir <- int(EventManager.MOTOR_DIR_UP)

	for {
		order := EventManager.Eventmanager_check_button_signal()
		if order.dir != 2 {
			//New order is detected
			chan_new_order <- order
		}

		chan_current_floor <- int(EventManager.Statemachine_set_current_floor())
		EventManager.Statemachine_set_button_lights()
		EventManager.Eventmanager_new_order_in_empty_queue()
		EventManager.Eventmanager_arrive_at_floor()
		EventManager.Eventmanager_door_timeout()
		EventManager.Eventmanager_orders_in_same_floor()
		time.Sleep(400 * time.Millisecond)
	}
}
