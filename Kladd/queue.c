#include "elev.h"
#include "statemachine.h"
#include <stdio.h>

int orders[3][4] = {{0,0,0,0},{0,0,0,0},{0,0,0,0}};


int check_orders_above(int current_floor){
	for (int i = current_floor + 1; i < 4; i++){
		if (orders[0][i] == 1 || orders[1][i] == 1 || orders[2][i] == 1){
			return 1;
		}
	}
	return 0;
}

int check_orders_below(int current_floor){ 
	for (int i = current_floor-1; i > -1; i--){
		if (orders[0][i] || orders[1][i] || orders[2][i]){
			return 1;
		}
	} 
	return 0;
}

    
int should_stop(dir_t dir, int current_floor) { 
	if(orders[2][current_floor]){
		return 1;
	}
	if (dir == UP){
		if(orders[0][current_floor] || !check_orders_above(current_floor)){
    		return 1;
		}
	}
	if (dir == DOWN){
		if(orders[1][current_floor] || !check_orders_below(current_floor)){
	    	return 1;
		}
	}
	if (dir == STILL){
		if(orders[0][current_floor] || orders[1][current_floor] || orders[2][current_floor]){
	    	return 1;
		}
	}
	return 0;
}
    
void delete_queue() { 
	int m;
	int n;
	for (m = 0; m < 3; m++) {
		for (n = 0; n < 4; n++) {
			orders[m][n] = 0;
		}
	}
}

void delete_orders(int current_floor) { 
	int i;
	for (i = 0; i < 3; i++) {
		if (orders[i][current_floor] != 0) {
			orders[i][current_floor] = 0;
		}
	} 
}

void queue_add_new_order(int floor, int button) {
	orders[button][floor] = 1;
}

int nr_of_orders_in_queue() {
	int m;
	int n;
	int nr_of_orders=0;
	for(m=0; m < 3; m++){
		for(n=0;n < 4; n++){
			if (orders[m][n]!=0) {
				nr_of_orders++;
			}
		}
	}
	return nr_of_orders;
}
