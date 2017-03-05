#ifndef __Q_H__
#define __Q_H__

int orders[3][4];


int check_orders_above(int current_floor);

int check_orders_below(int current_floor);

int should_stop(int dir, int floor);

void delete_queue(); 

void delete_orders(int current_floor); 

void queue_add_order(int floor, int button);


#endif