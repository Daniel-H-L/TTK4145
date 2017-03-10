package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

func Run_elevator(chan_state chan int, chan_dir chan int, chan_floor chan int, chan_order_executed chan int, chan_new_hw_order chan Button, chan_new_master_order chan Button, chan_set_lights chan [3][4]int) {
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
	go Statemachine_set_lights(chan_set_lights)

	timer := time.NewTimer(3 * time.Second)
	timer.Stop()

	for {
		select {
		case new_order := <-chan_new_master_order:
			fmt.Println("NEW ORDER: ", new_order)
			Statemachine_button_push(new_order, chan_dir, chan_state, timer)
			time.Sleep(50 * time.Millisecond)

		case floor := <-chan_floor_sensor:
			ElevatorFloor = floor

			//fmt.Println("NEW FLOOR: ", floor)
			EventManager.Elevator_set_floor_indicator(floor)
			Statemachine_arrived_at_floor(floor, chan_order_executed, timer)
			//time.Sleep(50 * time.Millisecond)

		case <-timer.C:
			Statemachine_door_time_out(chan_dir, chan_state)

			// case executed := <-chan_order_executed:
			// 	fmt.Println("executed", executed)
		}
	}
}
