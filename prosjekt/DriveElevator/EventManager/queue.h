#ifndef __Q_H__
#define __Q_H__

#include "statemachine.h"


int orders[3][4];


int check_orders_above(int current_floor);

int check_orders_below(int current_floor);

int should_stop(int dir, int floor);

void delete_q(); 

void delete_orders(int current_floor); 





#endif
