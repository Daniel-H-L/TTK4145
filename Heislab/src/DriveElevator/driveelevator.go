package DriveElevator

import (
	"./EventManager"
	"fmt"
	//"time"
)

func Run_elevator(chan_state chan int, chan_dir chan int, chan_floor chan int, chan_order_executed chan int, chan_new_hw_order chan Button, chan_new_master_order chan Button, chan_set_lights chan [][]int) {
	fmt.Println("Run elev...")
	if EventManager.Elevator_init() != 1 {
		fmt.Println("Uanble to initialize elevator hardware... \n")
		return
	}

	EventManager.Elevator_set_motor_dir(1)
	for {
		if EventManager.Elevator_get_floor_sensor_signal() != -1 {
			break
		}
		fmt.Println("Floor sensor is -1...")
	}
	EventManager.Elevator_set_motor_dir(0)

	chan_timer := make(chan bool, 1)
	chan_floor_sensor := make(chan int, 1)
	go Internal_queue_poll_buttons(chan_new_hw_order)
	go Internal_queue_poll_floor_sensors(chan_floor_sensor)
	go timer(chan_timer)

	for {
		select {
		case new_order := <-chan_new_master_order:
			fmt.Println("NEW ORDER: ", new_order)
			Statemachine_button_push(new_order, chan_dir)

		case floor := <-chan_floor_sensor:
			ElevatorFloor = floor
			HwState.Floor = floor

			fmt.Println("NEW FLOOR: ", floor)
			EventManager.Elevator_set_floor_indicator(floor)
			Statemachine_arrived_at_floor(floor, chan_order_executed)
			chan_state <- ElevatorState

		case timeout := <-chan_timer:
			if timeout {
				Statemachine_door_time_out(chan_dir)
			}
		case set_lights := <-chan_set_lights:
			Statemachine_set_lights(set_lights)
		}
	}
}
