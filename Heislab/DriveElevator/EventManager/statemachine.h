#ifndef __STATEMACHINE_H__
#define __STATEMACHINE_H__

typedef enum {
    	IDLE,
    	MOVING,
    	STOP 
} state_t;

typedef enum {
    	DOWN,
    	STILL,
    	UP
} dir_t;


state_t     state;

dir_t       dir;

int         current_floor;


int set_current_floor();

void set_button_lights();

void set_state_and_dir(state_t s, dir_t d);

int arrived_floor(int floor);


#endif
