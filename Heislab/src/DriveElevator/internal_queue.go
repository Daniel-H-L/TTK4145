package DriveElevator

import (
	"./EventManager"
)

var Internal_queue [3][4]int

type Button struct {
	dir   int
	floor int
}

func Internal_queue_add_new_order(floor int, dir int) {
	Internal_queue[dir][floor] = 1
}

func Internal_queue_delete_order(current_floor int) {
	for i := 0; i < 3; i++ {
		Internal_queue[i][current_floor] = 0
	}
}

func Internal_queue_delete_queue() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			Internal_queue[i][j] = 0
		}
	}
}

func Internal_queue_check_orders_above(current_floor int) int {
	for i := current_floor + 1; i < 4; i++{
		if (Internal_queue[0][i] == 1 || Internal_queue[1][i] == 1 || Internal_queue[2][i] == 1){
			return 1
		}
	}
	return 0
}

func Internal_queue_check_orders_below(current_floor int) int { 
	for i := current_floor-1; i > -1; i--{
		if (Internal_queue[0][i] == 1 || Internal_queue[1][i]== 1 || Internal_queue[2][i] == 1){
			return 1
		}
	} 
	return 0
}

    
func Internal_queue_should_stop(dir int, current_floor int) int { 
	if(Internal_queue[2][current_floor] == 1){
		return 1
	}
	if (dir == UP){
		if(Internal_queue[0][current_floor] == 1 || Internal_queue_check_orders_above(current_floor) == 0){
    		return 1
		}
	}
	if (dir == DOWN){
		if(Internal_queue[1][current_floor] == 1 || Internal_queue_check_orders_below(current_floor) == 0){
	    	return 1
		}
	}
	if (dir == STOP){
		if(Internal_queue[0][current_floor] == 1 || Internal_queue[1][current_floor] == 1 || Internal_queue[2][current_floor] == 1){
	    	return 1
		}
	}
	return 0
}


func Internal_queue_choose_dir() int {
	switch MotorDir {
	case UP:
		if Internal_queue_check_orders_above(ElevatorFloor) == 1 {
			return UP
		} else if Internal_queue_check_orders_below(ElevatorFloor) == 1 {
			return DOWN
		} else {
			return STOP
		}
	case DOWN:
		if Internal_queue_check_orders_below(ElevatorFloor) == 1 {
			return DOWN
		} else if Internal_queue_check_orders_above(ElevatorFloor) == 1 {
			return UP
		} else {
			return STOP
		}

	case STOP:
		if Internal_queue_check_orders_above(ElevatorFloor) == 1 {
			return UP
		} else if Internal_queue_check_orders_below(ElevatorFloor) == 1 {
			return DOWN
		} else {
			return STOP
		}
	default:
		return STOP
	}

}

func Internal_queue_poll_buttons(chan_new_hw_order chan Button) {

	var buttonStatus [3][4]bool
	for{
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if EventManager.Elevator_get_button_signal(i, j) == 1 && !buttonStatus[i][j] {
					button := Button{i, j}
					buttonStatus[i][j] = true
					chan_new_hw_order <- button
				} else if EventManager.Elevator_get_button_signal(i, j) == 0{
					buttonStatus[i][j] = false
				}
			}
		}
	}
}

func Internal_queue_poll_floor_sensors(chan_floor_sensor chan int) {
	var prev_floor int
	for {
		floor := EventManager.Elevator_get_floor_sensor_signal()
		if floor != -1 && prev_floor != floor {
			chan_floor_sensor <- floor
			prev_floor = floor 
		}
	}
}