#ifndef HARDWARESIMULATOR_H
#define HARDWARESIMULATOR_H

// hw_init() Creates a thread, that processes all gui events and updates the state 
// that can be checked with other functions in the interface.
void hw_init(); 

// Checking inputs
int hw_button1Status();
int hw_button2Status();

// Controlling the display
void hw_setDisplayNumber(int i); // Set the number to display 
void hw_on();  // The on and off functionality is a bit artificial compared to
void hw_off(); // the spec, since we need to "simulate" a display that is off. 
// hw_off() is not even used as it is.

#endif
