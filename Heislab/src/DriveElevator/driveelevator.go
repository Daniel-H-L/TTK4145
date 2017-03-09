package DriveElevator

import (
	"./EventManager"
	"fmt"
	//"time"
)

// func Driveelevator_check_events(chan_state chan int, chan_dir chan int, chan_floor chan int, chan_order_executed chan int, chan_new_order chan Orders) {
// 	select{
// 	case order := EventManager.Eventmanager_check_button_signal():
// 		if order.dir != -1 {
// 			//New order is detected
// 			chan_new_order <- order
// 		}
// 	case floor := Statemachine_send_deleted_order():
// 		if floor != -1 {
// 			chan_order_executed <- floor
// 		}
// 	case current_floor := EventManager.Statemachine_set_current_floor():
// 		if current_floor != -1 {
// 			chan_floor <- int(current_floor)
// 		}
// 	}
// }


func Run_elevator(chan_state chan int, chan_dir chan int, chan_floor chan int, chan_order_executed chan int, chan_new_hw_order chan Button, chan_new_master_order chan Button) { 
	fmt.Println("Run elev...")
	if EventManager.Elevator_init() != 1 {
		fmt.Println("Uanble to initialize elevator hardware... \n")
		return
	}

	EventManager.Elevator_set_motor_dir(1)
	for {
		if EventManager.Elevator_get_floor_sensor_signal() != -1  {
			break
		}
		fmt.Println("Floor sensor is -1...")
	}
	EventManager.Elevator_set_motor_dir(0)
/*
	EventManager.Statemachine_set_state_and_dir(EventManager.IDLE, EventManager.MOTOR_DIR_STOP)
	chan_state <- int(EventManager.IDLE)
	chan_dir <- int(EventManager.MOTOR_DIR_STOP)
*/
	chan_timer := make(chan bool, 1)
	chan_floor_sensor := make(chan int, 1)
	go Internal_queue_poll_buttons(chan_new_hw_order)
	go Internal_queue_poll_floor_sensors(chan_floor_sensor)
	go timer(chan_timer)

	for {
		select {
			case new_order := <- chan_new_master_order:
				fmt.Println("NEW ORDER: ", new_order)
				Statemachine_button_push(new_order)

			case floor := <- chan_floor_sensor:
				ElevatorFloor = floor

				fmt.Println("NEW FLOOR: ", floor)
				EventManager.Elevator_set_floor_indicator(floor)
				Statemachine_arrived_at_floor(floor)

			case timeout := <- chan_timer:
				if timeout {
					Statemachine_door_time_out()	
				}
		}
	}
}
