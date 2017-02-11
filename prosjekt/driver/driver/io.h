//
//  io.h
//  driver
//
//  Created by Mia Olea Vettestad on 10.02.2017.
//  Copyright © 2017 Mia Olea Vettestad. All rights reserved.
//

#ifndef io_h
#define io_h

#include <stdio.h>
#pragma once

// Returns 0 on init failure
int io_init(void);

void io_set_bit(int channel);
void io_clear_bit(int channel);

int io_read_bit(int channel);

int io_read_analog(int channel);
void io_write_analog(int channel, int value);


#endif /* io_h */
