#include "queue.h"
#include "elev.h"
#include "timer.h" 
#include <stdio.h>
#include "statemachine.h"

state_t     state;

dir_t       dir;

int         current_floor;

int set_current_floor() { 
	if(elev_get_floor_sensor_signal() == 0) {
		elev_set_floor_indicator(0);
		current_floor = 0;
		return 1;
	}
	if (elev_get_floor_sensor_signal() == 1) {
		elev_set_floor_indicator(1);
		current_floor = 1;
		return 2;
	}
	if (elev_get_floor_sensor_signal() == 2) {
		elev_set_floor_indicator(2);
		current_floor = 2;
		return 3;
	}
	if (elev_get_floor_sensor_signal() == 3) {
		elev_set_floor_indicator(3);
		current_floor = 3;
		return 4;
	}

return 0;
 
}

void set_button_lights() {
    int i;
    for (i = 0; i < 4; i++) {
        if (i != 3) {
            if (orders[0][i] == 1) {
                elev_set_button_lamp(BUTTON_CALL_UP, i, 1);
            }
            else if (orders[0][i] == 0) {
                elev_set_button_lamp(BUTTON_CALL_UP, i, 0);
            }
        }
        if (i != 0) {
            if(orders[1][i] == 1) {
                elev_set_button_lamp(BUTTON_CALL_DOWN, i, 1);
            }
            else if (orders[1][i] == 0) {
                elev_set_button_lamp(BUTTON_CALL_DOWN, i, 0);
            }
            
        }
        
        if (orders[2][i] == 1) {
            elev_set_button_lamp(BUTTON_COMMAND, i, 1);
        }
        else if (orders[2][i] == 0) {
            elev_set_button_lamp(BUTTON_COMMAND, i, 0);
        }
    }
}

void set_state_and_dir(state_t s, dir_t d){ 
	if (s == STOP) {
		elev_set_motor_direction(DIRN_STOP);
		state = STOP;
	}
	else if (s == IDLE){
		elev_set_motor_direction(DIRN_STOP);
		state = IDLE;
		dir = STILL;
	}
	else if (s == MOVING){
		if (d == UP) {
		    elev_set_motor_direction(DIRN_UP);
		    dir = UP;
		    state = MOVING;
		}
		else if (d == DOWN){
		    elev_set_motor_direction(DIRN_DOWN);
		    dir = DOWN;
		    state = MOVING;
		}
	}
}

int arrived_floor(int new_floor) { 
    current_floor = new_floor;
	if (should_stop(dir, current_floor)) {
		if (dir == UP){
			set_state_and_dir(STOP, UP); 
			timer_start();
			delete_orders(current_floor);
			elev_set_door_open_lamp(1);
			return 1;
		}
		else if (dir == DOWN){
			set_state_and_dir(STOP, DOWN); 
			timer_start();
			delete_orders(current_floor);
			elev_set_door_open_lamp(1);
			return 1;
		}	
		else if (dir == STILL){
			set_state_and_dir(STOP, STILL); 
			timer_start();
			delete_orders(current_floor);
			elev_set_door_open_lamp(1);
			return 1;
		}
			
	}	
	return 0;
}


