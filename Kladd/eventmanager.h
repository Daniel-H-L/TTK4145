#ifndef __EVENTMANAGER_H__
#define __EVENTMANAGER_H__

typedef struct Orders_s{
    int floor;
    int dir;
}Orders_t;

int new_order_in_empty_queue();

void arrive_at_floor();

void orders_in_same_floor();

int door_time_out();

Orders_t check_button_signal();

int stop_mechanical_reason();


#endif