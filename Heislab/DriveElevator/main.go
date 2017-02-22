package main

import (
	"./EventManager"
	"./driver"
	"fmt"
	"time"
)

func main() {
	if !driver.Elevator_init() {
		fmt.Println("Uanble to initialize elevator hardware... \n")
		return
	}
	for driver.Elevator_get_floor_signal() == -1 {
		EventManager.Statemachine_set_state_and_dir(EventManager.MOVING, driver.MOTOR_DIR_UP)
	}
	EventManager.Statemachine_set_state_and_dir(EventManager.IDLE, driver.MOTOR_STOP)
	for {
		EventManager.Eventmanager_check_button_signal()
		EventManager.Statemachine_set_current_floor()
		EventManager.Statemachine_set_button_lights()
		EventManager.Eventmanager_new_order_in_empty_queue()
		EventManager.Statemachine_arrive_at_floor()
		EvenManager.Eventmanager_door_timeout()
		EventManager.Eventmanager_ordeers_in_same_floor()
		time.Sleep(400 * time.Millisecond)
	}
}
