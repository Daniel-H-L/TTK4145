#include "queue.h"
#include "elev.h"
#include "timer.h" 
#include "eventmanager.h"
#include "statemachine.h"
#include <stdio.h>


int new_order_in_empty_queue(){ 
    if (state == IDLE) {
        if(check_orders_above(current_floor) == 1){
            set_state_and_dir(MOVING, UP);
        }
            
        else if(check_orders_below(current_floor) == 1){
            set_state_and_dir(MOVING, DOWN);
        }
    }
    return 0;
}
	

void arrive_at_floor(){
    static int prevFloor;
    int floor = elev_get_floor_sensor_signal(); 
    if(prevFloor != floor  &&  floor != -1){
        arrived_floor(floor);
    }
    prevFloor = floor;
}

void orders_in_same_floor() {
	int floor = elev_get_floor_sensor_signal();
	if(check_button_signal() == 1 && current_floor == floor) {
		arrived_floor(floor);
	}
}

int door_time_out(){
    if (state == STOP && timer_isTimeOut() == 1) {
	elev_set_door_open_lamp(0);
	timer_stop(); 
        if (dir == UP) {
            
            if(check_orders_above(current_floor) == 1){
                set_state_and_dir(MOVING, UP);
                return 1;
            }
        
            else if(check_orders_below(current_floor) == 1){
                set_state_and_dir(MOVING, DOWN);
                return 1;
            }
            
            else{
                set_state_and_dir(IDLE, STILL);
                return 1;
            }
        }
        
        else if (dir == DOWN){
            
            if(check_orders_below(current_floor) == 1){
                set_state_and_dir(MOVING, DOWN);
                return 1;
            }
            
            else if(check_orders_above(current_floor) == 1){
                set_state_and_dir(MOVING, UP);
                return 1;
            }
            
            else{
                set_state_and_dir(IDLE, STILL);
                return 1;
            }
        }
		else if (dir == STILL){
                set_state_and_dir(IDLE, STILL);
				return 1;
		}
        
    }
return 0;
}

int check_button_signal() {
    int i;
    for (i = 0; i < 4; i++) {
        if (i != 3) {
            if (elev_get_button_signal(BUTTON_CALL_UP, i) == 1) {
                orders[0][i] = 1;
		return 1; 
            }
        }
        if (i != 0) {
            if (elev_get_button_signal(BUTTON_CALL_DOWN, i) == 1) {
                orders[1][i] = 1;
		return 2; 
            }
        }
        
        if (elev_get_button_signal(BUTTON_COMMAND, i) == 1) {
		orders[2][i] = 1;
		return 3; 
		
        }
    }
	return 0;     
}

